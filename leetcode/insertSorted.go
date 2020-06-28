package main

import (
	"fmt"
)

func main() {
	arr := []int{1,2,3,4,5,6}
	//i := binarySearch(arr, 3)
	fmt.Println(binarySearch(arr, 3))
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

func binarySearch(s []int, key int) int {
	begin := 0
	end := len(s) - 1
	for {
		mid := (end - begin) / 2
		if s[mid] == key {
			return mid
		} else if s[mid] < key {
			begin = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}