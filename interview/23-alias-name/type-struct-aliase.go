package main

import "fmt"

/*
	定义类型别名, 等价与原类型，类型别名与原类型拥有相同的方法
*/
type User struct {}
//新的类型
type User1 User
//别名
type User2 = User

func (u User1) m1() {
	fmt.Println("m1")
}

func (u User) m2() {
	fmt.Println("m2")
}
func main() {
	var u1 User1
	var u2 User2
	u1.m1()
	u2.m2()
}
