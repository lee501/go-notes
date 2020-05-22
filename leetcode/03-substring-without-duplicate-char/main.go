package main

import "fmt"

/*
	给定一个字符串，找出不不含有重复字符的最长字串的长度

	"pwwkew"的最长字串为"wke" 长度为3
*/

func SubString(s string) int {
	charMap := make(map[byte]bool, len(s))
	curlen := 0
	for i:=0; i< len(s); i++ {
		if charMap[s[i]] == true {
			loop := i - curlen
			for loop < i {
				delete(charMap, s[loop])
				curlen--
				if s[loop] == s[i] {
					break
				}
				loop++
			}
		}
		charMap[s[i]] = true
		curlen++
	}
	return  curlen
}

func main() {
	s := "pwwkew"
	fmt.Println(SubString(s))
}