package roman

import (
	"fmt"
	"testing"
)

func TestIntToRoman(t *testing.T) {
	i := 3999
	m := IntToRoman(i)
	fmt.Println(m)
}
