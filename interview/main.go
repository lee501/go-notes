package main

import "fmt"

func main() {
	dict := map[string]string{
		"a": "this is a",
	}
	if v, ok := dict["a"]; ok {
		fmt.Println(v)
	}
}
