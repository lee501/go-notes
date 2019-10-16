package atoi

import (
	"fmt"
	"testing"
)

func TestMAtoi(t *testing.T) {
	str := "  -123 bs"
	fmt.Println(MAtoi(str))
}
