package closure_func

import (
	"fmt"
	"time"
)

//示例
//返回一个函数
func outer(x int) func(int) int {
	//return func(y int) int {
	//	return x + y
	//}
	var a =  func(y int) int {
		return x + y
	}
	return a
}

//for range使用闭包的注意的坑
//在没有将变量 v 的拷贝值传进匿名函数之前，只能获取最后一次循环的值
func Itera(s []string) {
	//s := []string{"a", "b", "c"}
	for _, v := range s {
		go func(v string) {
			fmt.Println(v)
		}(v)
	}
	time.Sleep(time.Second)
}

//函数列表处理函数列表
func processFunc() []func() {
	var s []func()
	for i := 0; i < 3; i++ {
		x := i  // 赋值变量
		s = append(s, func() {
			fmt.Println(x)
		})
	}
	return s
}

func processFunc1() (s []func()) {
	for i := 0; i < 3; i++ {
		x := i  // 赋值变量
		s = append(s, func() {
			fmt.Println(x)
		})
	}
	return s
}

//延迟调用
func DeferPrint() {
	x, y := 1, 2

	defer func(a int) {
		fmt.Println(a, y) //此处y为闭包引用外部的y
	}(x) //此处赋值x为1

	x += 100
	y += 100
	fmt.Println(x, y)
}

