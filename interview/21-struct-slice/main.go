package main

import "fmt"

type Foo struct {
	Bar string
}

func main() {
	s := []Foo{Foo{"a"}, {"b"}, {"c"}}
	s1 := make([]*Foo, len(s))

	for i, value := range s {
		s1[i] = &value
	}
	fmt.Println(s)
	fmt.Println(s1) //[0xc000086030 0xc000086030 0xc000086030] c的地址
}
