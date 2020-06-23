package main

import "fmt"

func main() {
	arr := []int{3,2,5,1}
	insertSort(arr)
	fmt.Println(arr)
}

func insertSort(array []int) {
	if len(array) < 2 {
		return
	}
	for i:=1; i < len(array); i++ {
		pre := i - 1
		key := array[i]
		for pre >= 0 && array[pre] > key {
			array[pre+1], array[pre] = array[pre], array[pre+1]
			pre--
		}
	}
}
