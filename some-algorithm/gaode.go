package main

import (
	"fmt"
	"sort"
	"strconv"
)

/*
	// 评测题目: 实现一个功能，这个功能是从一个字符串，并找出其中对应的整形数值，
	// 排序去重后输出，以及出现的计数
	// #example input addcce13ddds12ddcd1122211111fff12dcd12dcddddd11
	// #output value count
	// # 1122211111 1
	// # 13 1
	// # 12 3
*/

func processString(s string) []int {
	if s == "" {
		return nil
	}
	list := make([]int, 0)
	res := ""
	for i := 0; i < len(s)-1; i++ {
		switch s[i] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			res += string(s[i])
		default:
			if res != "" {
				k,_ := strconv.Atoi(res)
				list = append(list, k)
				res = ""
			}
		}
	}
	return list
}


func countInt(data []int) [][]int {
	// "1122211111"
	res := make([][]int, 0)
	so := make([]int, 0)
	m := make(map[int]int)
	for _, v := range data {
		if k, ok := m[v]; ok {
			m[v] = k+1
			continue
		}
		m[v] = 1
	}
	fmt.Println(m)
	for k, v := range m {
		so = append(so, k, v)
		res = append(res, so)
		so = []int{}
	}
	return res
}

func main() {
	s := "addcce13ddds12ddcd1122211111fff12dcd12dcddddd11"
	re := processString(s)
	sort.Ints(re)
	fmt.Println(re)
	fmt.Println(countInt(re))
}