package regular

import (
	"fmt"
	"testing"
)

func TestIsMatch(t *testing.T) {
	s := "abcd"
	p := "a.*"
	fmt.Println(IsMatch(s, p))
}
