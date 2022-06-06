package main

import (
	"fmt"
)

func main() {
	var s1 interface{} = "123"
	var s2 []byte = []byte{'4', '5', '6'}
	fmt.Println(fmt.Sprintf("%s%s", s1, s2))
}
