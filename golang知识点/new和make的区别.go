/*
new和make都在堆上分配内存

new 函数分配内存，make 函数初始化

new(T): 适用于值类型和结构体struct， 返回一个指针 &T{}

make(T): 只用于内建的引用类型slice、map、channal
*/

// 代码示例
package main 

import "fmt"

func main() {
  // p == nil; with len and cap 0
  p := new([]int)
  fmt.Println(p)

  //  v is initialed with len 10, cap 50
  v := make([]int, 10, 20)
  fmt.Println(v)

  /*********Output****************
        &[]
        [0 0 0 0 0 0 0 0 0 0]
  *********************************/
  (*p)[0] = 1  //会报错panic: runtime error: index out of range, because p is a nil pointer, with len and cap 0
  v[1] = 18
}
