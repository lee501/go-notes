package main

import (
	"fmt"
	"time"
)

/*
	slice删除元素
	a = append(a[:i], a[i+1:]...)
// or
a = a[:i+copy(a[i:], a[i+1:])]
*/

func main() {
	s := []string{"a", "b", "c", "d"}
	begin := time.Now()
	s = append(s[:1], s[2:]...)
	end := time.Now()
	fmt.Println(end.UnixNano() - begin.UnixNano()) //0
	m := []string{"a", "b", "c", "d"}
	begin = time.Now()
	var s2 []string
	for _, v := range m {
		if v != "b" {
			s2 = append(s2, v)
		}
	}
	end = time.Now()
	fmt.Println(end.UnixNano() - begin.UnixNano()) //1000~ 7000
}
