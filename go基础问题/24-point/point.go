package main

import "fmt"

/*
	考点： 指针
*/
func incre(p *int) int {
	*p++
	return *p
}
func main() {
	v := 1
	//传递引用地址
	incre(&v)
	fmt.Println(v)
}
