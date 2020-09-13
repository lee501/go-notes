package main

import "fmt"
//funcOption 将func(int) int 的函数转换成 funcOption 类型

type funcOption func(int) int

func sum(i int) int {
  return i * 2
}

func main() {
   option := funcOption(sum)
   m := option(2)
   fmt.Println(m) 
}
