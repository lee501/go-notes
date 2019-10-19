package main

import (
	"fmt"
	"regexp"
)

/*
	给定一个数组和目标值，找出数组中和为目标值的两个数
	nums = [2, 7, 11, 15] , target = 9
	返回格式下标值[0, 1]
*/
func TweSum(nums []int, target int) []int {
	hash := make(map[int]int)
	//遍历将数组的下标和值放入map处理
	for i, value := range nums {
		//查询另一个值是否在map中
		if j, ok := hash[target - value]; ok {
			return []int{i, j}
		}
		hash[value] = i
	}
	return nil
}

func main()  {
	s := "349"
	//字节转int, -'0'操作
	fmt.Println(int(s[0]-'0'))
	m := " "
	fmt.Println(len(m))
	re := regexp.MustCompile("")
	fmt.Println(re.Match([]byte("")))
}
