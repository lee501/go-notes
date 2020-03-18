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
	_ = append(s[:1], s[2:]...)
	end := time.Now()
	fmt.Println(s)
	fmt.Println(end.UnixNano() - begin.UnixNano()) //0
	m := []string{"a", "b", "c", "d"}
	begin = time.Now()
	//开辟新的内存存放切片，防止更改问题
	var s2 []string
	for _, v := range m {
		if v != "b" {
			s2 = append(s2, v)
		}
	}
	end = time.Now()
	fmt.Println(end.UnixNano() - begin.UnixNano()) //1000~ 7000
}
