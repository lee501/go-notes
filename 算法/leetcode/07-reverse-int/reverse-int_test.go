package reverse

import "testing"

func TestReverse(t *testing.T) {
	a := 123
	b := 2147483648
	a = Reverse(a)
	b = Reverse(b)
	if a != 321 {
		t.Error("reverse error, expected result is 321")
	}
	if b != 0 {
		t.Error("reverse error, expected result is 0")
	}
}
