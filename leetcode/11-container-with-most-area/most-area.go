package container

func MaxArea(height []int) int {
	i, j := 0, len(height) - 1
	max := 0

	//计算最大面积
	for i < j {
		l1, l2 := height[i], height[j]
		//最小值
		h := min(l1, l2)
		area := h * (j - i)
		if max < area {
			max = area
		}
		//移动游标
		if l1 < l2 {
			i++
		} else {
			j--
		}
	}
	return max
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}