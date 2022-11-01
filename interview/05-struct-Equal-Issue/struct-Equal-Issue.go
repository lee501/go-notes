package main

import "fmt"

/*
	考点：
		关于结构体相等问题
		1. 相同类型的结构体才能进行比较，属性类型和顺序必须都要相同
		2. 成员都为值类型才能比较
*/
func main() {
	s1 := struct {
		age  int
		name string
	}{age: 11, name: "lee"}

	s2 := struct {
		age  int
		name string
	}{age: 11, name: "lee"}

	//s1和s2可以比较
	if s1 == s2 {
		fmt.Println("s1 == s2")
	}

	n1 := struct {
		age int
		m   map[string]string
	}{age: 12, m: map[string]string{"a": "1"}}

	n2 := struct {
		age int
		m   map[string]string
	}{age: 12, m: map[string]string{"a": "1"}}

	fmt.Println(n1, n2)
	//n1 和 n2无法比较，不能通过编译
	//if n1 == n2 {
	//
	//}
}
