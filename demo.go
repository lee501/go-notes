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
func t() {
	switch {
	case false:
		fmt.Println("The integer was <= 4")
		fallthrough
	case true:
		fmt.Println("The integer was <= 5")
		fallthrough
	case false:
		fmt.Println("The integer was <= 6")
		fallthrough
	case true:
		fmt.Println("The integer was <= 7")
		fallthrough
	default:
		fmt.Println("default case")
	}
}
