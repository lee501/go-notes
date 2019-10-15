package zigzag

import (
	"fmt"
	"testing"
)

func TestZigZagConvert(t *testing.T) {
	str := "abcdef"
	num := 3
	res := ZigZagConvert(str, num)
	fmt.Println(res)
}