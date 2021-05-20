package main

func main() {
//	test slidingWindow
}

func slidingWindow(data []int, k int) []int {
	var result []int
	if len(data) == 0 {
		return result
	}
	if len(data) <= k {
		result = append(result, getMax(data))
		return result
	}
	var windowSlice []int
	index := 0
	for i := k; i <= len(data); i ++ {
		windowSlice = data[index:i]
		result = append(result, getMax(windowSlice))
		index++
	}
	return result
}

func getMax(data []int) int {
	max := data[0]
	for _, v :=  range data {
		if v > max {
			max = v
		}
	}
	return max
}