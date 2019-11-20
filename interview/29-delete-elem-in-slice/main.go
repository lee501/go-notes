package main

import "fmt"

/*
	slice删除元素的坑
	copy复制会比等号复制慢。但是copy复制为值复制，改变原切片的值不会影响新切片。
	而等号复制为指针复制，改变原切片或新切片都会对另一个产生影响。
*/
func main() {
	nums := []int{1,2,3,4}
	k := 2
	//res := append(nums[:k], nums[k+1:]...)
	//fmt.Println(res) // [1 2 4]
	//fmt.Println(nums) // [1 2 4 4]

	//正确处理方式
	temp := make([]int, len(nums[:k]))
	copy(temp, nums[:k])
	temp = append(temp, nums[k+1:]...)
	fmt.Println(temp)
	fmt.Println(nums)
}
