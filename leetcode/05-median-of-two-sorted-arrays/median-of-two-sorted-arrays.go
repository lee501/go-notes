package median

import "fmt"

/*
	合并两个已经排序的切片，返回合并后切片的中位数
	例子：
		nums1 = [1, 3]
		nums2 = [2]
		The median is 2.0
	例子2:
		nums1 = [1, 3]
		nums2 = [2, 4]
		The median is 2.5
*/
func FindMedianSortedArrays(nums1, nums2 []int) float64 {
	nums := combine(nums1, nums2)
	return medianOf(nums)
}

func combine(left, right []int) []int {
	m := len(left)
	n := len(right)
	i, j := 0, 0
	res := make([]int, m + n)
	//新数组长度为两个数组和， 交叉比较, 考虑边界问题
	for k := 0; k < m + n; k ++ {
		if i == m || (i < m && j < n && left[i] >= right[j]) {
			res[k] = right[j]
			j++
			continue
		}
		if j == n || (i < m && j < n && left[i] < right[j]) {
			res[k] = left[i]
			i++
		}
		fmt.Println(i, j)
		fmt.Println(res)
	}
	return res
}

func medianOf(nums []int) float64 {
	l := len(nums)
	if l == 0 {
		panic("切片的长度为0，无法求解中位数。")
	}
	if l % 2 == 0 {
		return float64(nums[l/2] + nums[l/2 - 1]) / 2.0
	}
	return float64(nums[1/2])
}
