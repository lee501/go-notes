package http_simple_proxy

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func NewMultiHostsRevsrseProxy(targets []*url.URL) *httputil.ReverseProxy {
	director := func(req *http.Request) {
		target := targets[rand.Int()%2]
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.URL.Path = target.Path
	}
	return &httputil.ReverseProxy{Director: director}
}

func Print() {
	num := rand.Int() % 2
	fmt.Println(num)
}
