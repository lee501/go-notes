package main

import (
	"fmt"
	"reflect"
)

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
	pe := people{"lee"}
	fmt.Println(reflect.ValueOf(pe).FieldByName("name"))
	str := "abc"
	fmt.Println(str[0])
	for _,v := range str{
		fmt.Println(v)
	}
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