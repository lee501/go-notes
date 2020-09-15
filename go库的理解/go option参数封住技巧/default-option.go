package main

import "time"

//grpc设置默认参数代码示例
type dialOption struct {
	insecure bool
	timeout  time.Duration
}

type DialOption interface {
	apply(*dialOption)
}

//属性为方法的结构体， 方法的参数为dialOption, 与继承的接口apply方法的参数保持一致
type funcDialOption struct {
	f func(*dialOption)
}


func (fdo *funcDialOption) apply(do *dialOption) {
	fdo.f(do)
}

//构造函数NewFunDialOption，接收一个函数作为参数
func newFunDialOption(f func(*dialOption)) *funcDialOption {
	return &funcDialOption{
		f: f,
	}
}

//newFuncDialOption的调用
//开启不安去调用
func WithInsecure() DialOption {
	return newFunDialOption(func(option *dialOption) {
		option.insecure = true
	})
}
//设置超时时间
func WithTimeout(d time.Duration) DialOption {
	return newFunDialOption(func(option *dialOption) {
		option.timeout = d
	})
}

/*
	来体验一下这里的精妙设计：
			1. 首先对于每一个字段，提供一个方法来设置其对应的值。由于每个方法返回的类型都是 DialOption ，
			从而确保了 grpc.DialContext 方法可用可选参数，因为类型都是一致的；
			2. 返回的真实类型是 *funcDialOption ，但是它实现了接口 DialOption，这增加了扩展性。
*/