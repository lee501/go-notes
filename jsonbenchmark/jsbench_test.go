package jsonbenchmark

import (
	"testing"

	"github.com/bytedance/sonic"
	"github.com/goccy/go-json"
)

type AgentService struct {
	ServiceName    string
	Version        string
	ServiceId      string
	Address        string
	Port           int
	Metadata       map[string]string
	ConnectTimeOut int
	ConnectType    string
	ReadTimeOut    int
	WriteTimeOut   int
	Protocol       string
	Balance        string
	Idcs           string
	Converter      string
	Retry          int
}

var obj = AgentService{
	ServiceName:    "kaleidoscope_api",
	Version:        "1517558949087295000_1298498081",
	ServiceId:      "kaleidoscope.com_v1.2",
	Address:        "127.0.0.1",
	Port:           80,
	Metadata:       map[string]string{},
	ConnectTimeOut: 1000,
	ConnectType:    "LONG",
	ReadTimeOut:    1000,
	WriteTimeOut:   1000,
	Protocol:       "HTTP",
	Balance:        "Random",
	Idcs:           "hu,hd,hn",
	Converter:      "json",
	Retry:          3,
}

func BenchmarkSonic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := sonic.Marshal(obj)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkGoJson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(obj)
		if err != nil {
			b.Error(err)
		}
	}
}
