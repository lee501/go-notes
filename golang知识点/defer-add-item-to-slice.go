package main

import "fmt"

type MSlice []int

func NewSlice() MSlice{
	return make(MSlice, 0)
}

func (s *MSlice) Add(elem int) *MSlice {
	*s = append(*s, elem)
	fmt.Print(elem)
	return s
}

func main() {
	s := NewSlice()
	defer s.Add(1).Add(2).Add(3)
	s.Add(4)
}

type A struct {
	X int `json:"x"`
}
