package main

import "fmt"

//标准组合关系
/*
type People struct {
	name string
	age int
}

type Teacher struct {
	people People
	classroom string
}

func main() {
	teacher := Teacher{People{"lee", 17}, "三年二班"}
	fmt.Println(teacher)
}
*/

//使用匿名属性完成Go语言中的继承
type People struct {
	name string
	age  int
}

type Teacher struct {
	People  //匿名组合
	classroom string //班级
}

func main() {
	teacher := Teacher{People{"smallming", 17}, "302教室"}
	fmt.Println(teacher.classroom, teacher.age, teacher.name)
}
