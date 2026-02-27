package atoi

import (
	"math"
	"strings"
)

/*
	example
		Input: "42"
		Output: 42

		Input: "   -42"
		Output: -42

		Input: "4193 with words"
		Output: 4193

		Input: "words and 987"
		Output: 0

	超过32-bit signed integer range: [−2^31, 2^31 − 1]， 返回INT_MAX (2^31 − 1) or INT_MIN (−2^31)
*/
func MAtoi(s string) int {
	sign, str := processString(s)
	return convertInt(sign, str)
}

func processString(s string) (sign int, str string) {
	//处理字符串首尾空格
	s = strings.TrimSpace(s)
	if s == "" {
		return
	}
	//先处理字符串首部
	switch s[0] {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		sign, str = 1, s
	case '+':
		sign, str = 1, s[1:]
	case '-':
		sign, str = -1, s[1:]
	default:
		str = ""
		return
	}
	//遍历查看字符串中的数字项
	for i, ch := range str {
		if ch < '0' || ch > '9' {
			str = str[:i]
			break
		}
	}
	return
}

func convertInt(sign int, str string) int {
	res := 0
	for _, i := range str {
		//字节转int
		res = res * 10 + int(i-'0')
		switch {
		case sign == 1 && res > math.MaxInt32:
			return math.MinInt32
		case sign == -1 && res*sign < math.MinInt32:
			return math.MinInt32
		}
	}
	return sign*res
}
