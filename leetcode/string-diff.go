package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "abcdeab"
	fmt.Println(strings.Count(s, "a"))
	fmt.Println(isUniqueString(s))
}

func isUniqueString(s string) bool{
	if strings.Count(s, "") > 256 {
		return false
	}
	for _, v := range s {
		if v > 127 {
			return false
		}
		if strings.Count(s, string(v)) > 1 {
			return false
		}
	}
	return true
}