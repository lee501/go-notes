package longestOnes

import "testing"

func TestLongestOnes(t *testing.T) {
	array := []int{1,1,0,0,0,0,1,1,1}
	K := 3
	re := LongestOnes(array, K)
	if re != 6 {
		t.Error("error, expected value is 6")
	}
}
