package _5_day

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
	fmt.Println(end.UnixNano() - begin.UnixNano())
	m := []string{"a", "b", "c", "d"}
	begin = time.Now()
	re := make([]string, len(m)-1)
	re = append(re, m[:1]...)
	re = append(re, m[2:]...)
	end = time.Now()
	fmt.Println(end.UnixNano() - begin.UnixNano())
}
