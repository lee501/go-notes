package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	s := "[1, 2]"
	var r []int
	json.Unmarshal([]byte(s), &r)
	fmt.Println(r)
}