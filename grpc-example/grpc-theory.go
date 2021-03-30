package grpc_example

import (
	"context"
	"fmt"
	"time"
)

//grpc dial options set
type dialOption struct {
	//摘取部分参数
	block       bool
	insecure    bool
	timeout     time.Duration
}
//DialOption how to config option
type DialOption interface {
	apply(option *dialOption)
}
//funcDialOption wraps a function that modifies dialOptions
type funcDialOption struct {
	f func(*dialOption)
}

func newFuncDialOption(f func(*dialOption)) *funcDialOption {
	return &funcDialOption{
		f: f,
	}
}

func (fdo *funcDialOption) apply(option *dialOption) {
	fdo.f(option)
}

func WithBlockDemo() DialOption {
	return newFuncDialOption(func(option *dialOption) {
		option.block = true
	})
}

func WithInsecureDemo() DialOption {
	return newFuncDialOption(func(option *dialOption) {
		option.insecure = true
	})
}

func WithTimeoutDemo(d time.Duration) DialOption {
	return newFuncDialOption(func(option *dialOption) {
		option.timeout = d
	})
}

//对外暴露借口
func DialContextDemo(ctx context.Context, target string, ops ...DialOption) {
	//
	dialOption := &dialOption{}
	for _, opt := range ops {
		opt.apply(dialOption)
	}
	fmt.Printf("--------%#v\n-----------", dialOption)
}
func Demo() {
	ctx, cancel := context.WithTimeout(context.Background(), 10* time.Minute)
	defer func() {
		cancel()
	}()
	target := "localhost:3000"
	DialContextDemo(ctx, target, WithBlockDemo(), WithInsecureDemo(), WithTimeoutDemo(10*time.Minute))
}