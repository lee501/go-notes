### 流式代理

#### 用法
```go
package main

import (
    "log"
    "net/http"
)

func main() {
    // 创建各种代理处理器
    httpProxy := NewAdvancedStreamProxy()
    wsProxy := NewWebsocketProxy("localhost:8081")
    sseProxy := &SSEProxy{backend: "http://localhost:8082"}
    
    // 设置路由
    mux := http.NewServeMux()
    mux.Handle("/api/", httpProxy)
    mux.Handle("/ws/", wsProxy)
    mux.Handle("/events/", sseProxy)
    
    // 健康检查
    mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("OK"))
    })
    
    // 启动服务器
    server := &http.Server{
        Addr:         ":8080",
        Handler:      mux,
        ReadTimeout:  30 * time.Second,
        WriteTimeout: 30 * time.Second,
        IdleTimeout:  120 * time.Second,
    }
    
    log.Println("Starting stream proxy on :8080")
    if err := server.ListenAndServe(); err != nil {
        log.Fatal("Server failed:", err)
    }
}
```

#### 限流中间件
```go
  // 限流中间件
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
```