package main

import "fmt"

func main() {
	arr := []int{5, 4, 7, 8, 2, 4, 3, 6, 7, 2}
	QuickSort(arr, 0, len(arr) - 1)
	fmt.Println(arr)
}

func QuickSort(array []int, begin, end int){
	if len(array) < 2{
		return
	}
	temp := array[begin]
	i, j := begin, end
	for {
		for i < j {
			if array[j] < temp {
				array[i], array[j] = array[j], array[i]
				break
			}
			j--
		}
		fmt.Println(array, i, j) //2, 4, 7, 8, 2, 4, 3, 6, 7, 5
		for i < j {
			if array[i] >= temp {
				array[j],array[i] = array[i], array[j]
				fmt.Println(array)
				break
			}
			i++
		}
		if i >= j {
			array[i] = temp
			break
		}
	}

	if i -1 > begin {
		QuickSort(array, begin, i-1)
	}

	if j + 1 < end {
		QuickSort(array, j+1, end)
	}
}
