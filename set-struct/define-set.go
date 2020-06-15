package main

import "fmt"

func main() {
	foo()
}

var a = 10
func foo() {
	fmt.Println(a)
	fmt.Println(a)
}