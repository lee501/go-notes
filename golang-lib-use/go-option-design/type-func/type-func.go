package main

import (
	"fmt"
    "time"
)
//funcOption 将func(int) int 的函数转换成 funcOption 类型

type funcOption func(int) int

func sum(i int) int {
  return i * 2
}

func main() {
   option := funcOption(sum)
   m := option(2)
   fmt.Println(m)
   fd := SetInsecure()
   di := new(dialOption)
   fd.apply(di)
   fmt.Println(di)
}

//定一个func
type funcDial func(*dialOption)

func (fd funcDial) apply(do *dialOption) {
    fd(do)
}

//对外的接口
func SetInsecure() DialOption {
    return funcDial(func(d *dialOption){
        d.insecure = true
    })
}

type dialOption struct {
    insecure bool
    timeout  time.Duration
}

type DialOption interface {
    apply(*dialOption)
}