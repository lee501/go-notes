package main

import "fmt"

func main() {
	//str := []string{"a","b","c"} []string 不可以转换成[]interface{}
	multi("a", "b")
}

func multi(s ...interface{}) {
	for _, v := range s {
		fmt.Println(v)
	}
}
