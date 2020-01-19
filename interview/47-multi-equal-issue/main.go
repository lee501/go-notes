package main

import "fmt"

/*
	多重赋值的顺序
		1. 先计算等号左边的索引表达式和取址表达式， 接着计算等号右边的表达式
		2. 赋值
*/

func assignIssue() {
	var k = 1
	var s = []int{1, 2}
	k, s[k] = 0, 3
	fmt.Println(s[1])
}

func main() {
	assignIssue()
	//s := []int{1,2,3}
	var k int = 9
	//range空切片不会遍历
	for k = range []int{} {
		fmt.Println(k+1)
	}
	//range nil指针数组
	fmt.Println((*[3]int)(nil))

	for k = range (*[3]int)(nil) {
		fmt.Println(k) //输出0 1 2
	}
	fmt.Println(k) //输出为2

	//for循环
	for k = 0; k < 3; k++ {

	}
	fmt.Println(k) //输出为3
}
