package http_simple_proxy

import (
	"net/http"
	"net/url"
	"testing"
)

func TestPrint(t *testing.T) {
	for i := 0; i < 20; i++ {
		Print()
	}
}

func TestNewMultiHostsRevsrseProxy(t *testing.T) {
	proxy := NewMultiHostsRevsrseProxy([]*url.URL{
		{
			Scheme: "http",
			Host:   "localhost:9091",
		},
		{
			Scheme: "http",
			Host:   "localhost:9092",
		},
	})
	http.ListenAndServe(":9090", proxy)
}
