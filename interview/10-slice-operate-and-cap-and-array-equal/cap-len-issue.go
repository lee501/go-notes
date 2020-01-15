package main

import "fmt"

/*
	切片操作符 [low,high]
		规则0 <= low <= high <= cap(原切片)
*/
func main() {
	s := make([]int, 3, 9)
	fmt.Println(len(s), cap(s))

	s1 := s[4:8]
	fmt.Println(len(s1), cap(s1))
}
