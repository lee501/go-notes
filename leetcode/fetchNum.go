package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["a"] = 1
	a := m["a"]
	fmt.Println(a)
	nums := []int{1,1,2,3,4,2,3,4,5}
	fmt.Println(fetchNum(nums))
}

func fetchNum(nums []int) int {
	l := len(nums)
	for i := 0; i < l; i++ {
		flag := false
		for j := 0; j < l; j++ {
			if nums[i] == nums[j] && i != j {
				flag = true
			}
		}
		if !flag {
			return nums[i]
		}
	}
	return -1
}
