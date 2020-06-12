package main

import "fmt"

type Foo struct {
	Bar string
}

func main() {
	s := []Foo{{"a"}, {"b"}, {"c"}}
	s1 := make([]*Foo, len(s))

	for i, value := range s {
		s1[i] = &value
	}
	fmt.Println(s)
	fmt.Println(s1[1]) //[0xc000086030 0xc000086030 0xc000086030] c的地址

	 //map的无序性
	var m = map[string]int{
		"A": 21,
		"B": 22,
		"C": 23,
	}
	counter := 0
	for k, v := range m {
		if counter == 0 {
			delete(m, "A")
		}
		counter++
		fmt.Println(k, v)
	}
	fmt.Println("counter is ", counter)
}
