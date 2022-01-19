package main

import "fmt"

type Name string

const (
	NA Name = "1"
	YA Name = "2"
)

func ceshi(n string)  {
	switch Name(n) {
	case NA:
		fmt.Println("aNa")
	case YA:
		fmt.Println(2)
	}
}

func main()  {
	s := "1"
	ceshi(s)
}
