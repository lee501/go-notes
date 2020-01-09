package longestOnes

/*
	滑动窗口问题
		最大连续1的个数
			给定一个由若干 0 和 1 组成的数组 A，我们最多可以将 K 个值从 0 变成 1 。
			返回仅包含 1 的最长（连续）子数组的长度

		输入：A = [1,1,1,0,0,0,1,1,1,1,0], K = 2
		输出：6
		解释：
		[1,1,1,0,0,1,1,1,1,1,1]
		粗体数字从 0 翻转到 1，最长的子数组长度为 6。

	解题思路：
		1，维持窗口内0的个数<=K
		2，我们定义指针l,r分别表示窗口左右下标，移动r，当A[r]＝＝0的时候我们增加0的个数记录sum，分两种情况
			A,sum>K    这个时候需要移动左指针，让0的个数减1
			B,sum<=K  无需处理，继续移动右指针
*/
func LongestOnes(array []int, K int) int {
	if K == 0 || len(array) < 1 {
		return 0
	}
	l := 0
	sum := 0
	max := 0

	for r := 0; r < len(array); r++ {
		if array[r]==0{
			sum++
			if sum>K{
				for array[l]!=0{
					l++
				}
				l++
				sum--
			}
		}

		if r - l + 1 > max {
			max = r -l + 1
		}
	}

	return max
}
