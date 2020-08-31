package main

import (
	"fmt"
	"sort"
)

//sort满足以下接口
/*
	type Interface interface {
		Len() int

		Less(i, j int) int

		Swap(i, j int)
	}
*/
type Man struct {
	Age int
}

type manSlice []Man

func (ms manSlice) Len() int {
	return len(ms)
}

func (ms manSlice) Less(i, j int) bool {
	return ms[i].Age < ms[j].Age
}

func (ms manSlice) Swap(i, j int) {
	ms[i], ms[j] = ms[j], ms[i]
}

func main() {
	manslice := []Man{{5}, {4}, {3}}
	sort.Sort(manSlice(manslice))
	fmt.Println(manslice)
}
