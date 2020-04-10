package context

import "time"

//Context接口， context中的方法是协程安全的，在父routine中创建context,可以传递给任意数量的routine，让他们同时访问
type Context interface {
	//返回一个超时时间，routine获得超时时间后，可以对某些io操作设定超时时间
	Deadline() (deadline time.Time, ok bool)
	//返回一个channel，当该context被取消时，该channel会被关闭，同时对应使用该context的routine也应该结束返回
	Done() <-chan struct{}
	Err() error
	//让routine共享一些数据，获得数据是协程安全的
	Value(key interface{}) interface{}
}
  