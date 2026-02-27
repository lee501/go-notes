package advancedStreamProxy

import (
	"context"
	"io"
	"net"
	"net/http"
	"time"
)

type AdvancedStreamProxy struct {
	transport http.RoundTripper
	timeout   time.Duration
}

func NewAdvancedStreamProxy() *AdvancedStreamProxy {
	return &AdvancedStreamProxy{
		transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second, //TLS 握手的超时时限
			ExpectContinueTimeout: 1 * time.Second,
			DisableCompression:    true, //不自动声明 Accept-Encoding: gzip ，也不做自动解压。这避免了解压带来的缓冲和延迟
		},
		timeout: time.Second * 10,
	}
}

func (p *AdvancedStreamProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), p.timeout)
	defer cancel()

	req := r.Clone(ctx)
	//逐跳（hop-by-hop）”头
	RemoveHopHeaders(req.Header)

	//发送请求
	resp, err := p.transport.RoundTrip(req)
	if err != nil {
		http.Error(w, "Gateway error: "+err.Error(), http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()
	// 复制响应头部
	CopyHeaders(w.Header(), resp.Header)
	w.WriteHeader(resp.StatusCode)

	// 关键：流式复制响应体
	//io.Copy(w, resp.Body)
	p.streamCopy(w, resp.Body)
}

func (p *AdvancedStreamProxy) streamCopy(dst io.Writer, src io.Reader) {
	buf := make([]byte, 1024)
	flusher, hasFlusher := dst.(http.Flusher)
	for {
		n, err := src.Read(buf)
		if n > 0 {
			if _, writeErr := dst.Write(buf[:n]); writeErr != nil {
				break
			}
			// 如果是可刷新的writer，立即刷新
			if hasFlusher {
				flusher.Flush()
			}
		}
		if err != nil {
			break
		}
	}
}

/*
  - Connection ：声明本连接特有的头；代理应移除它及其列出的头名。

- Proxy-Connection ：非标准，旧客户端用于代理连接控制，不能转发。
- Keep-Alive ：连接保活信息，仅对当前连接有意义。
- Proxy-Authenticate / Proxy-Authorization ：用于客户端与代理的认证，不是后端的认证。
- TE ：仅允许 trailers ，其他值不应转发；通常直接删除。
- Trailer ：声明将出现在消息尾部的头，仅与当前传输有关。
- Transfer-Encoding ：消息分帧（如 chunked ）属于传输层，后端应自行决定；代理不应强灌此头。
- Upgrade ：协议升级（如 WebSocket）是连接级别行为，需要特殊处理，不应盲目转发。
*/
func RemoveHopHeaders(header http.Header) {
	// 移除 Hop-by-hop 头部
	hopHeaders := []string{
		"Connection",
		"Proxy-Connection",
		"Keep-Alive",
		"Proxy-Authenticate",
		"Proxy-Authorization",
		"Te",
		"Trailer",
		"Transfer-Encoding",
		"Upgrade",
	}

	for _, h := range hopHeaders {
		header.Del(h)
	}
}

// 工具函数
func CopyHeaders(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}
