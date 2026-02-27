package main

import "sync"

func main() {

}
func mergeSort(nums []int) []int {
	n := len(nums)
	if n < 2 {
		return nums
	}
	key := n / 2
	left := mergeSort(nums[0:key])
	right := mergeSort(nums[key:])
	return merge(left, right)
}

// 并发版本
func concurrentMergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2

	var left, right []int
	var wg sync.WaitGroup
	wg.Add(2)

	// 并发处理左半部分
	go func() {
		defer wg.Done()
		left = concurrentMergeSort(arr[:mid])
	}()

	// 并发处理右半部分
	go func() {
		defer wg.Done()
		right = concurrentMergeSort(arr[mid:])
	}()

	// 等待两个goroutine完成
	wg.Wait()

	// 合并结果
	return merge(left, right)
}

// 合并两个有序数组
func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}
