package main

func StrMulti(str1, str2 string) string {
	str1 = Reverse(str1)
	str2 = Reverse(str2)
	re := make([]byte, len(str1)+len(str2))

	for i := 0; i < len(str1); i++ {
		for j := 0; j < len(str2); j++ {
			t := (str1[i] - '0') * (str2[j] - '0')
			re[i+j] += t
		}
	}

	var plus byte = 0
	result := ""
	for i := 0; i < len(re); i++ {
		if re[i] == 0 {
			break
		}
		temp := re[i] + plus
		plus = 0
		if temp > 9 {
			plus = temp / 10
			result += string(temp - plus*10 + '0')
		} else {
			result += string(temp + '0')
		}

	}
	return Reverse(result)
}

//翻转字符串
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

/*
func main() {
	str1 := "1234567899"
	str2 := "9987654321"
	s := StrMulti(str1,str2)
	fmt.Println(s)
}
*/
