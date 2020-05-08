package main

import "fmt"

type people struct {
	name string
}
func main() {
	//arr := []int{1,2,3,4}
	//fmt.Println(arr[:2])
	//fmt.Println(arr[2:])
	//a :=[...]int{1,2,3}
	//m := reflect.TypeOf(a)
	//fmt.Printf("%v", m)
	var p *people
	fmt.Println(p == nil)
}

type M interface {
	N
}

type N interface {
	Check()
}

func (p *people) Check() {
	fmt.Println("interface")
}