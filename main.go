package main

import "fmt"

func main()  {
	s := "浙江省杭州市"
	storune := []rune(s)
	fmt.Println(string(storune[:2]))
}