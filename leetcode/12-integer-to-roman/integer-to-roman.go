package roman

//表驱动法：转换数字成罗马
func IntToRoman(num int) string {
	d := [4][]string {
		{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"},
		{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"},
		{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"},
		{"", "M", "MM", "MMM"},
	}
	return d[3][num/1000] + d[2][num/100%10] + d[1][num/10%10] + d[0][num%10]
}
