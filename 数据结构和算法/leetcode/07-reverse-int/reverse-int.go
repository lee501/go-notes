package reverse

import "math"

/*
	example:
		123 => 321
		-123 => 321
*/
func Reverse(x int) int {
	sign := 1
	res := 0
	if x < 0 {
		sign = -1
		x = x * sign
	}
	for x > 0 {
		temp := x % 10
		res = res * 10 + temp
		x = x / 10
	}
	res = res * sign
	if res > math.MaxInt32 || res < math.MinInt32 {
		res = 0
	}
	return res
}
