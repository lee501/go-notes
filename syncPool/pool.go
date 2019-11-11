package main

import (
	"fmt"
	"reflect"
	"sync"
	"time"
)

/*
	go临时对象池
		sync.Pool是一组临时对象的集合， Pool是协程安全的
		Pool用于存储那些被分配了但是没有被使用，而未来可能会使用的值，减少内存压力
*/
//结构体放入pool中的demo
type buffer []byte

type pp struct {
	buf buffer
	arg interface{}
	value reflect.Value
	//fmt fmt
	reordered bool
	goodArgNum bool
	panicking bool
	erroring bool
}

//pp对象池
var ppFree = sync.Pool{
	New: func() interface{} {
		return new(pp)
	},
}

//分配新的pp或从对象池中拿一个
func newPrinter() *pp {
	p := ppFree.Get().(*pp)
	p.panicking = false
	p.erroring = false
	//p.fmt.init(&p.buf)
	return p
}


//简单的示例
//[]byte对象池
var bytePool = sync.Pool{
	New: func() interface{} {
		b := make([]byte, 1024)
		return &b
	},
}

func main() {
	a := time.Now().Unix()
	//不使用对象池
	for i := 0; i < 1000000000; i++ {
		obj := make([]byte, 1024)
		_ = obj
	}
	b := time.Now().Unix()
	//使用对象池
	for i := 0; i < 1000000000; i++ {
		obj := bytePool.Get().(*[]byte)
		_ = obj
		bytePool.Put(obj)
	}
	c := time.Now().Unix()
	fmt.Println("without pool ", b - a, "s")
	fmt.Println("with    pool ", c - b, "s")
}
