package palinfrome

import "strconv"

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	s := convert2String(x)
	for i, j := 0, len(s)-1; i < j; {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func convert2String(x int) string {
	return strconv.Itoa(x)
}
