package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
	声明为变量时，并分配内存空间, 在进行struct中的变量赋值时，可直接使用。
	声明为指针时，并未进行赋值，也没有分配内存空间
*/
var c Man

type Man struct {
	Name string
}

func main() {
	var a *Man
	//指针必须指定内存地址，否则//invalid memory address or nil pointer dereference
	a = &Man{}
	var b Man
	if a == nil {
		fmt.Printf("777: \n", 1)
	}
	if &b == nil {
		fmt.Printf("888: ", &b)
	}
	a.Name = "lee"
	b.Name = "lee"
	fmt.Println(a)
	fmt.Println(b)

	c.Name = "lee"
	fmt.Println()

	s, _ := M()
	fmt.Println(s)
}

func M() (string, error) {
	b := []byte{'"', '{', '}', '"'}
	str := string(b)
	return strconv.Unquote(strings.Replace(strconv.Quote(str), `\\u`, `\u`, -1))
}
