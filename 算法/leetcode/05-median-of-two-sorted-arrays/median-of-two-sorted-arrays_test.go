package median

import (
	"fmt"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	nums1 := []int{1, 3, 6}
	nums2 := []int{3, 4, 5}
	re := FindMedianSortedArrays(nums1, nums2)
	if re != 3.5 {
		fmt.Println(re)
		t.Error("result is error, expected result is 3.5")
	}
}
