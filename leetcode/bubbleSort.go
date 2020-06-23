package main

import "fmt"

func main() {
	var array []int = []int{1,9, 8,3,6}
	bubbleSort(array)
	fmt.Println(array)
}

func bubbleSort(array []int) {
	if len(array) < 2 {
		return
	}
	for i := 0; i < len(array) - 1; i++ {
		for j := 0; j < len(array) - 1 - i; j++ {
			if array[j] > array[j +1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
}