package main

import "fmt"

func main() {
	m := make(map[string]int)
	apply(m)
	fmt.Print(m)
}

func apply(m map[string]int) {
	m["a"] = 1
}