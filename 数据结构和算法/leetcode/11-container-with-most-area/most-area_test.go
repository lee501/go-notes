package container

import (
	"fmt"
	"testing"
)

func TestMaxArea(t *testing.T) {
	array := []int{4, 6, 8, 7, 3}
	re := MaxArea(array)
	fmt.Println(re)
}
