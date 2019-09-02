package main

import (
	"fmt"
	"strings"
)

type SplitFunc func(data string) (result []string)

type Mystruct struct{
	split SplitFunc
	Iscalled bool
}
func main()  {
	//s := "浙江省杭州市"
	//storune := []rune(s)
	//fmt.Println(string(storune[:2]))
	//input := bufio.NewScanner(os.Stdin)
	a := false
	suffer := NewSuffer(a)
	suffer.Callback()

	m := NewMystruct(false)
	m.Split(test)
	m.split("test")
}

func NewMystruct(called bool) Mystruct {
	return Mystruct{
		split: SplitString,
		Iscalled: called,
	}
}

func SplitString(data string) (result []string){
	return strings.Split(data, ",")
}

func (m *Mystruct) Split(split SplitFunc) {
	if m.Iscalled {
		panic("Split called after Scan")
	}
	m.split = split
}

type Suffer struct {
	check bool
}

func (s *Suffer) Callback() {
	if s.check {
		panic("test the next action")
	}
	fmt.Println(111)
}

func NewSuffer(a bool) *Suffer {
	return &Suffer{
		check: a,
	}
}

func test(s string) ([]string) {
	var str []string
	fmt.Println(s)
	return append(str, s)
}
