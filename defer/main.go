package main

import "fmt"

//defer 常用的场景是关闭链接
//多个defer 以栈结构执行，先产生后执行

/*defer跟return 同时存在的时候，return理解成两条执行语句（非原子指令）：1。给返回值赋值， 2。跳出函数
	defer和return执行顺序
	*return赋值给返回值
	*执行defer
	*跳出函数
*/

//没有定义返回值的接收变量
func demo() int {
	i := 0
	defer func() {
		i = i + 2
	}()
	return i //结果返回的是0
}

//定义返回值的接收变量

func demo2() (i int){
	i = 1
	defer func() {
		i = i + 3
	}()
	return i //返回的是4, 原因是函数执行前已经声明了内存空间给i
}

//defer 和panic： panic不是立即停止程序(os.Exit(0)),defer还是在panic前执行的.

func main() {
	for {
		fmt.Println(2)
		defer fmt.Println(1)
	}
}