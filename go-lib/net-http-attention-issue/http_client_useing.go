package net_http_attention_issue

import (
	"net"
	"net/http"
	"time"
)

//1. 默认的HttpClient不包含请求超时时间，如果你使用http.Get(url)或者&Client{}, 这将会使用http.DefaultClient，这个结构体内no timeout
func NewClient() *http.Client {
	return &http.Client{
		Timeout: 10 * time.Second,
	}
}

//2. go的Transport 可理解为连接池中的连接, 不要使用默认Transport，增加MaxIdleConnsPerHost
var DefaultTransport http.RoundTripper = &http.Transport{
	Proxy: nil,
	DialContext: (&net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second, //KeepAlive是tcp探活的时间间隔，并不是我们HTTP连接复用的 Keep-Alive
	}).DialContext,
	ForceAttemptHTTP2:     true,
	MaxIdleConns:          100,              //Http Client连接池有100个连接
	IdleConnTimeout:       90 * time.Second, //每个连接默认的空闲时间90s(90s内有请求过来，可以复用该连接)
	TLSHandshakeTimeout:   10 * time.Second,
	ExpectContinueTimeout: 1 * time.Second,
}
