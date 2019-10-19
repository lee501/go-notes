package regular

/*
	Given an input string (s) and a pattern (p), implement regular expression matching with support for '.' and ''. Implement regular expression matching with support for '.' and ''.

		'.' Matches any single character.
		'*' Matches zero or more of the preceding element.

	notes:
		s could be empty and contains only lowercase letters a-z.
		p could be empty and contains only lowercase letters a-z, and characters like . or *.

	1. ""空字符串与后置位为*星号("a*")或空字符串("")样式匹配
	2.  re := regexp.MustCompile("")   fmt.Println(re.Match([]byte("")))

定义一个二维的 DP 数组，其中 dp[i][j] 表示 s[:i] 和 p[:j] 是否 match
dp[i][j] = dp[i - 1][j - 1], if p[j - 1] != '*' && (s[i - 1] == p[j - 1] || p[j - 1] == '.');
dp[i][j] = dp[i][j - 2], if p[j - 1] == '*' and the pattern repeats for 0 time;
dp[i][j] = dp[i - 1][j] && (s[i - 1] == p[j - 2] || p[j - 2] == '.'), if p[j - 1] == '*' and the pattern repeats for at least 1 time

*/
func IsMatch(s, p string) bool {
	slen, plen := len(s), len(p)
	//二维数组, 默认每个元素都是false
	b := make([][]bool, slen+1)
	for i := range b {
		b[i] = make([]bool, plen+1)
	}
	//b[i][j]为s[:i], p[:j]是否匹配
	//当s和p的长度都为0的时候，两个空字符串是匹配的
	b[0][0] = true
	//当s的长度为0， p为"a*b*"时匹配
	for j :=1; j < plen && b[0][j-1]; j += 2 {
		if p[j] == '*' {
			b[0][j+1] = true
		}
	}
	//字符串非空的情况
	for i := 0; i < slen; i++ {
		for j := 0; j < plen; j++ {
			if p[j] == '.' || p[j] == s[i] {
				/* p[j] 与 s[i] 可以匹配上，所以，只要前面匹配，这里就能匹配上 */
				b[i+1][j+1] = b[i][j]
			} else if p[j] == '*' {
				//无法匹配
				if p[j-1] != s[i] && p[j-1] != '.' {
					//*号前不匹配，匹配到s[:i+1], p[:j-1]
					b[i+1][j+1] = b[i+1][j-1]
				} else {
					/*
					 p[j] 与 s[i] 匹配上
					 p[j-1;j+1] 作为 "x*", 可以有三种解释
						""||"x"||"xxx..."
					 */
					b[i+1][j+1] = b[i+1][j-1] || b[i+1][j] || b[i][j+1]
				}
			}
		}
	}
	return b[slen][plen]
}
