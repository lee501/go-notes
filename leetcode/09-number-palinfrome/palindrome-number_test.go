package palinfrome

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	x := 1
	re := IsPalindrome(x)
	if !re {
		t.Error("check number is error, expected result is true")
	}
	x = 1221
	fmt.Println(IsPalindrome(x))
}
