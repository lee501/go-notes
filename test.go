package main

import (
	"fmt"
	"reflect"
)

func main() {
	//arr := []int{1,2,3,4}
	//fmt.Println(arr[:2])
	//fmt.Println(arr[2:])
	a :=[...]int{1,2,3}
	m := reflect.TypeOf(a)
	fmt.Printf("%v", m)
}
