package main

import (
	"encoding/json"
	"fmt"
)

/*
	1.channel为nil，读写都会阻塞
	2.结构体访问控制，结构体中的属性大写，外部才能访问
*/

func main() {
	chanIssue()
	jsonToStruct()
}

func chanIssue() {
	//未分配内存，ch为nil
	var ch chan int
	select {
	case v, ok := <-ch:
		println(v, ok)
	default:
		println("default")
	}
}

type People struct {
	Name string `json: "name"`
}

func jsonToStruct() {
	js := `{
		"name": "seekload"
	}`

	var p People
	err := json.Unmarshal([]byte(js), &p)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(p)
}
