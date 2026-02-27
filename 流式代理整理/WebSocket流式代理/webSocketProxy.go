package webSocketProxy

import (
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

// WebSocket 代理
type WebsocketProxy struct {
	backend string
}

func (p *WebsocketProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 升级到 WebSocket
	if !websocket.IsWebSocketUpgrade(r) {
		http.Error(w, "WebSocket upgrade required", http.StatusBadRequest)
		return
	}

	// 连接后端
	backendURL := *r.URL
	backendURL.Scheme = "ws"
	backendURL.Host = p.backend

	// 代理 WebSocket 连接
	p.proxyWebSocket(w, r, backendURL.String())
}

func (p *WebsocketProxy) proxyWebSocket(w http.ResponseWriter, r *http.Request, backend string) {
	// 连接到后端 WebSocket
	backendConn, _, err := websocket.DefaultDialer.Dial(backend, nil)
	if err != nil {
		http.Error(w, "Could not connect to backend", http.StatusBadGateway)
		return
	}
	defer backendConn.Close()

	// 升级客户端连接
	clientConn, err := websocket.Upgrade(w, r, nil, 1024, 1024)
	if err != nil {
		return
	}
	defer clientConn.Close()

	// 双向流式转发
	var wg sync.WaitGroup
	wg.Add(2)

	// 客户端 -> 后端
	go func() {
		defer wg.Done()
		for {
			msgType, data, err := clientConn.ReadMessage()
			if err != nil {
				break
			}
			if err := backendConn.WriteMessage(msgType, data); err != nil {
				break
			}
		}
	}()

	// 后端 -> 客户端
	go func() {
		defer wg.Done()
		for {
			msgType, data, err := backendConn.ReadMessage()
			if err != nil {
				break
			}
			if err := clientConn.WriteMessage(msgType, data); err != nil {
				break
			}
		}
	}()

	wg.Wait()
}
