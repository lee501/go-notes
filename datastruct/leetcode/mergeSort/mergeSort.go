package main
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
func merge(left, right []int) []int {
	newArr := make([]int, len(left)+len(right))
	i, j, index := 0, 0, 0
	for {
		if left[i] > right[j] {
			newArr[index] = right[j]
			index++
			j++
			if j == len(right) {
				copy(newArr[index:], left[i:])
				break
			}
		} else {
			newArr[index] = left[i]
			index++
			i++
			if i == len(left) {
				copy(newArr[index:], right[j:])
				break
			}
		}
	}
	return newArr
}
