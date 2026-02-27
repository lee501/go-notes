package SSEStreamProxy

import "net/http"

// Server-Sent Events
type SSEProxy struct {
	method  string
	backend string
}

func (p *SSEProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 设置 SSE 头部
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// 创建到后端的请求
	req, err := http.NewRequest("GET", p.backend+r.URL.Path, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 复制头部
	copyHeaders(req.Header, r.Header)

	// 发送请求
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	// 流式复制 SSE 数据
	flusher, _ := w.(http.Flusher)
	buf := make([]byte, 1024)

	for {
		n, err := resp.Body.Read(buf)
		if n > 0 {
			w.Write(buf[:n])
			flusher.Flush() // 立即刷新到客户端
		}
		if err != nil {
			break
		}
	}
}

// 工具函数
func copyHeaders(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}
