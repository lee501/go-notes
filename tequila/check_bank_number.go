package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func validateLuhn(bankId string) bool {
	source := strings.Split(bankId, "")
	checksum := 0
	double := false

	for i := len(source) - 1; i > -1; i-- {
		t, err := strconv.ParseInt(source[i], 10, 8)
		if err != nil {
			log.Println("validateLuhn", t, err)
			return false
		}
		n := int(t)
		if double {
			n = n * 2
		}
		double = !double
		if n >= 10 {
			n = n - 9
		}
		checksum += n
	}
	return (checksum % 10) == 0
}


func luhn(s string) bool {
	var t = [...]int{0, 2, 4, 6, 8, 1, 3, 5, 7, 9}
	odd := len(s) & 1
	var sum int
	for i, c := range s {
		if c < '0' || c > '9' {
			return false
		}
		if i&1 == odd {
			sum += t[c-'0']
		} else {
			sum += int(c - '0')
		}
	}
	return sum%10 == 0
}


func main() {
	s := "1234567812345678"
	fmt.Println(validateLuhn(s))
	fmt.Println(luhn(s))

	id := "139549203907235328"
	reg := "[1-9]{2}\\d{4}(19\\d{2}|20[0-2][0-9])((0[1-9])|(1[0-2]))(([0-2][1-9])|10|20|30|31)\\d{3}[0-9Xx]"
	re := regexp.MustCompile(reg)
	r := re.FindStringSubmatch(id)
	if len(r) > 0 {
		fmt.Println(len(r), r[0])
	}
}
