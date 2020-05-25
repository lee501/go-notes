package main

import "fmt"

/*
	考点：
		基于类型定义的方法必须在同一个包内
		错误示范
			func (i int) PrintInt() {
				fmt.Println(i)
			}
*/

type Mint int
func (i Mint) PrintInt() {
	fmt.Println(i)
}


/*
	考点：接口的*T和T的方法集
		*T类型的方法，只有*T可以接收
		T类型的方法，两者皆可接收
*/
type Person interface {
	Speak(string) string
}

type Student struct {}

func (s Student) Speak(think string) (talk string) {
	talk = think
	return
}

func main() {
	var i Mint = 1
	i.PrintInt()

	var p Person = &Student{}
	think := "hi"
	p.Speak(think)
}
