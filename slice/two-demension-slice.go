package main

import "fmt"

/*
	init 二维切片
*/
func main() {
	tslice := make([][]bool, 5)
	for i := range tslice {
		tslice[i] = make([]bool, 5)
	}
	//输出内容
	for i := range mslice {
		for j := range mslice[i] {
			fmt.Println(i, j, mslice[i][j])
		}
	}
}
