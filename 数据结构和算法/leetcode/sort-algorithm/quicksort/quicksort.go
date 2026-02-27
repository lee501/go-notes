package main

import (
	"fmt"
)

//func main() {
//	arr := []int{5, 4, 7, 8, 2, 4, 3, 6, 7, 2}
//	QuickSort(arr, 0, len(arr) - 1)
//	fmt.Println(arr)
//}
//
//func QuickSort(array []int, begin, end int){
//	if len(array) < 2{
//		return
//	}
//	temp := array[begin]
//	i, j := begin, end
//	for {
//		for i < j {
//			if array[j] < temp {
//				array[i], array[j] = array[j], array[i]
//				break
//			}
//			j--
//		}
//		fmt.Println(array, i, j) //2, 4, 7, 8, 2, 4, 3, 6, 7, 5
//		for i < j {
//			if array[i] >= temp {
//				array[j],array[i] = array[i], array[j]
//				fmt.Println(array)
//				break
//			}
//			i++
//		}
//		if i >= j {
//			array[i] = temp
//			break
//		}
//	}
//
//	if i -1 > begin {
//		QuickSort(array, begin, i-1)
//	}
//
//	if j + 1 < end {
//		QuickSort(array, j+1, end)
//	}
//}

func quickSort(s []int, begin, end int) {
	i := begin
	j := end
	temp := s[i]
	for i <= j {
		for s[i] < temp {
			i++
		}
		for s[j] > temp {
			j--
		}
		if i <= j {
			s[i], s[j] = s[j], s[i]
			i++
			j--
		}
	}
	if begin < j {
		quickSort(s, begin, j)
	}
	if end > i {
		quickSort(s, i, end)
	}
}

func main() {
	s := []int{4, 1, 6, 3, 5, 7, 2}
	quickSortV3(s)
	fmt.Println(s)
}

func quickSortV3(arr []int) {
	if len(arr) <= 1 {
		return
	}

	//选择中间元素作为基准值
	pivotIndex := len(arr) / 2
	pivotVal := arr[pivotIndex]
	//交换基准到最后位置
	arr[pivotIndex], arr[len(arr)-1] = arr[len(arr)-1], arr[pivotIndex]

	storeIndex := 0 //指向第一个大于基准的位置
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] < pivotVal {
			if i != storeIndex {
				arr[i], arr[storeIndex] = arr[storeIndex], arr[i]
			}
			storeIndex++
		}
	}
	//放回正确的位置
	arr[storeIndex], arr[len(arr)-1] = arr[len(arr)-1], arr[storeIndex]
	quickSortV3(arr[:storeIndex])
	quickSortV3(arr[storeIndex+1:])
}
