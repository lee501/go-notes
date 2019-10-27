package main

import "fmt"

/*
	考点：iota用法和String()方法
		类型定义String() 方法时候，fmt.Printf()、fmt.Print() 和 fmt.Println() 会自动使用 String() 方法
*/

type Direct int

const (
	East Direct = iota
	South
	West
	North
)

func (d Direct) String() string {
	return [...]string{"East", "South", "West", "North"}[d]
}

/*
	考点：
		map中的value本身不可以寻址 m["foo"].x = 1这种写法错误
*/

type Math struct {
	x, y int
}
//需要借助临时变量
var m = map[string]Math{
	"foo": Math{2, 3},
}
//地址方法
var n = map[string]*Math{"foo": &Math{1,2}}

func main() {
	fmt.Println(South)
	//借助临时变量
	temp := m["foo"]
	temp.x = 3
	m["foo"] =temp
	//value为指针
	n["foo"].x = 2
}
