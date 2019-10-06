package main

import (
	"fmt"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	s := "abcbdc"
	re := LongestPalindrome(s)
	fmt.Println(re)
}