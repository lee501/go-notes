package main

import "fmt"

type ShowItem int

const (
	First = iota
	Second
	Third
)

type Item struct {
	id ShowItem
}

func main() {
	item := Item{
		0,
	}
	if item.id == First {
		fmt.Println("1")
	}
}
