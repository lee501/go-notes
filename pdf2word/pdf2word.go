package main

import (
	"fmt"
)

func main() {
	cells := make([]string, 5)
	var b byte = 65
	for i, _ := range cells {
		cells[i] = string(b + byte(i))
		fmt.Println(cells[i])
	}
	fmt.Println(cells)
}
