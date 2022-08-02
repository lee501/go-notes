package main

import (
	"fmt"

	"github.com/bytedance/sonic"
)

func main() {
	m := struct {
		Name string
		Sex  string
	}{
		Name: "lee",
		Sex:  "man",
	}
	val, err := sonic.Marshal(&m)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(val))
}
