package baseStreamProxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"sync"
	"time"
)

/*
- 目标： http://backend/api （ u.Path="/api" ）
- 客户端请求： /v1/users （ req.URL.Path="/v1/users" ）
*/
// 基础反向代理: 使用httputil.ReverseProxy
type BaseStreamProxy struct {
	target *url.URL
	proxy  *httputil.ReverseProxy
}

func NewBaseStreamProxy(target string) *BaseStreamProxy {
	u, _ := url.Parse(target)
	proxy := &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			req.URL.Scheme = u.Scheme
			req.URL.Host = u.Host
			req.URL.Path = singleJoiningSlash(u.Path, req.URL.Path)
			req.Host = u.Host
			if u.RawQuery == "" || req.URL.RawQuery == "" {
				req.URL.RawQuery = u.RawQuery + req.URL.RawQuery
			} else {
				req.URL.RawQuery = u.RawQuery + "&" + req.URL.RawQuery
			}
			// 保持重要头部
			if _, ok := req.Header["User-Agent"]; !ok {
				req.Header.Set("User-Agent", "")
			}
		},
		// 关键：禁用缓冲，启用流式传输
		BufferPool:    newBufferPool(),
		FlushInterval: time.Millisecond * 100,
	}
	return &BaseStreamProxy{target: u, proxy: proxy}
}

func (bp *BaseStreamProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 设置流式响应头
	w.Header().Set("X-Stream-Proxy", "go-proxy")
	//调用反响代理
	bp.proxy.ServeHTTP(w, r)
}

// 自定义缓冲池，减少GC压力
type bufferPool struct {
	pool sync.Pool
}

// 自定义BufferPool，用于流式传输
func newBufferPool() *bufferPool {
	return &bufferPool{
		pool: sync.Pool{
			New: func() any {
				return make([]byte, 0, 1024)
			},
		},
	}
}

func (bp *bufferPool) Get() []byte {
	return bp.pool.Get().([]byte)
}

func (bp *bufferPool) Put(buf []byte) {
	bp.pool.Put(buf)
}

func singleJoiningSlash(a, b string) string {
	aslash := strings.HasSuffix(a, "/")
	bslash := strings.HasPrefix(b, "/")
	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		return a + "/" + b
	default:
		return a + b
	}
}

func mux() {
	httpProxy := NewBaseStreamProxy("127.0.0.1:90")

	// 设置路由
	mux := http.NewServeMux()
	limiter := NewRateLimiter(100)
	mux.Handle("/api/", Chain(httpProxy, limiter.Middleware))
}

type Middleware func(next http.Handler) http.Handler

func Chain(h http.Handler, m ...Middleware) http.Handler {
	for i := len(m) - 1; i >= 0; i-- {
		h = m[i](h)
	}
	return h
}

type RateLimiter struct {
	tokens chan struct{}
}

func NewRateLimiter(limit int) *RateLimiter {
	rl := &RateLimiter{
		tokens: make(chan struct{}, limit),
	}
	// 预先填充令牌
	for i := 0; i < limit; i++ {
		rl.tokens <- struct{}{}
	}
	return rl
}

func (rl *RateLimiter) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		select {
		case <-rl.tokens:
			defer func() { rl.tokens <- struct{}{} }()
			next.ServeHTTP(w, r)
		case <-time.After(100 * time.Millisecond):
			http.Error(w, "Too many requests", http.StatusTooManyRequests)
		}
	})
}

// 便捷函数：直接生成一个限流中间件
func RateLimit(limit int) Middleware {
	rl := NewRateLimiter(limit)
	return rl.Middleware
}
