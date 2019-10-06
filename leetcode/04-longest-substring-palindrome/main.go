package main

import (
	"math"
)

/*
	判断字符串的最长子回文
	"abcbdc"
*/

func LongestPalindrome(s string) string {
	l := len(s)
	if l <= 1 {
		return s
	}
	start, end := 0, 0
	for i := 0; i < l; i++ {
		len1 := cal(s, i, i) //中间回文为一个"abcbe"
		len2 := cal(s, i, i+1) //中间回文为两个
		len := int(math.Max(float64(len1), float64(len2)))
		//始终位置
		if len > end - start {
			//len=3， i=2
			start = i - (len- 1) / 2
			end = i + len/2
		}
	}
	return string(s[start:end+1])
}

func cal(s string, i, j int) int {
	l, r := i, j
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return r - l - 1
}
