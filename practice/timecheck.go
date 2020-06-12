package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t)
	m := t.Add(60 * time.Second)
	fmt.Println(m)
}