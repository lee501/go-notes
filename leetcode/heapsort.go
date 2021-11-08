package main

import "fmt"

/*
	堆排序：
		子节点i的父节点 (i-1)/2
		父节点i的左右子节点 2*i+1、2*i+2
*/
//构建堆 [4,6,8,5,9]
/*
4        4 		  9		  9
/\       /\		  /\	  /\
6 8 ->  9  8 ->  4  8    6  8 (构建大根堆) -> swap 首尾元素，重新构建剩下元素   8		5       6
/\      /\       /\      /\												   /\		/\		/\
5 9    5  6      5 6    5  4											  6   4 -> 6  4 -> 5  4
																		 /
																		5
*/
//len数组长度， i为最后一个非叶节点8 （len/2 -1）
func AdjustHeap(array []int, len, i int) {
	root := array[i]
	for l := 2*i + 1; l < len; l = l*2 + 1 {
		if l+1 < len && array[l] < array[l+1] {
			l++
		}
		if array[l] > root {
			array[i] = array[l]
			i = l
		}
	}
	array[i] = root
}

func HeapSort(array []int) {
	len := len(array)
	//从最后一个非叶子节点从下到上，从右到左调整
	for i := len/2 - 1; i >= 0; i-- {
		AdjustHeap(array, len, i)
	}
	//交换堆顶元素与末尾元素
	for j := len - 1; j > 0; j-- {
		swap(array, 0, j)

		AdjustHeap(array, j, 0)
	}
}

func swap(array []int, i, j int) {
	array[i], array[j] = array[j], array[i]
}

func main() {
	arr := []int{4, 6, 8, 5, 9}
	HeapSort(arr)
	fmt.Println(arr)
}
