package main

import "fmt"

/*
	指针和地址转换的练习
*/

type T struct {
	x int
	y *int
}

func main() {
	i := 20
	t := T{10, &i}

	//p为x的地址
	p := &t.x

	//对p取值后操纵++ --
	*p++
	*p--

	//更换结构体t中y的值
	t.y = p
	fmt.Println(*t.y)

	sliceIssue()
}

/*
	切片的截取运算
		[i:j] 默认j的取值 j < 原切片的cap， 当j缺省时，默认是原数组的长度
*/
func sliceIssue()  {
	x := make([]int, 2, 10)
	_ = x[6:10]
	//_ = x[6:] panic: runtime error: slice bounds out of rang
}