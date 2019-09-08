package main

import "fmt"

/*
	1.golang中的继承是通过struct匿名组合来实现的；
	2.多态是通过接口来实现的
	3.struct与interface之间是隐形继承的
	4.接口是对继承的一种补充
	5.继承主要解决代码的复用性和可维护性
	6.接口的价值设计各种规范，让其他struct实现这些方法，多态上多些
*/
//demo 猴子天生会爬树， 通过学习学会飞行和游泳
type Monkey struct {
	Name string
}

func (m *Monkey) Climbing() {
	fmt.Println(m.Name, "天生会爬树")
}

type LittleMOnkey struct {
	Monkey //匿名字段实现继承
}

func (lm LittleMOnkey) Flying(){
	fmt.Println(lm.Name, " 通过学习会飞翔...")
}

func (lm LittleMOnkey) Swimming(){
	fmt.Println(lm.Name, " 通过学习会游泳...")
}

func main() {
	lm := LittleMOnkey{Monkey{"悟空"}}
	lm.Climbing()
	lm.Flying()
	lm.Swimming()
}
