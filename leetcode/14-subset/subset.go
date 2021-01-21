package main

import (
	"fmt"
	"time"
)

func main() {
	nums := []int{1,2,3}
	b := time.Now()
	res := subsets(nums)
	fmt.Println(res, time.Since(b))
}

//func subsets(nums []int) [][]int {
//	var res [][]int
//	var temp []int
//	tracking(nums, &res, temp, 0)
//	return res
//}
//
//func tracking(nums []int, res *[][]int, temp []int, index int) {
//	co := make([]int, len(temp))
//	copy(co, temp)
//	*res = append(*res, co)
//	for i := index; i < len(nums); i++ {
//		temp = append(temp, nums[i])
//		tracking(nums, res, temp, i+1)
//		temp=temp[:len(temp)-1]
//	}
//}
func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	res = append(res, []int{})
	for i := len(nums)-1; i >= 0; i-- {
		length := len(res)
		for j := 0; j < length; j++ {
			tmp := make([]int, 0)
			tmp = append(tmp, nums[i])
			//这里的...的意思是把切片中的所有元素一个个地添加到tmp中
			tmp = append(tmp, res[j]...)
			res = append(res, tmp)
		}
	}
	return res
}
