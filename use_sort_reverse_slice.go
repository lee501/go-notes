package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strings"
)

type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

// Array 实现Interface接口
type Array []int

func (arr Array) Len() int {
	return len(arr)
}

func (arr Array) Less(i, j int) bool {
	return arr[i] < arr[j]
}

func (arr Array) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

// 匿名接口(anonymous interface)
type reverse struct {
	Interface
}

// 重写(override)
func (r reverse) Less(i, j int) bool {
	return r.Interface.Less(j, i)
}

// 构造reverse Interface
func Reverse(data Interface) Interface {
	return &reverse{data}
}

//func main() {
//	arr := Array{1, 2, 3}
//	rarr := Reverse(arr)
//	fmt.Println(arr.Less(0,1))
//	fmt.Println(rarr.Less(0,1))
//
//	var t Base
//	t = &Test{}
//	fmt.Println(t.Name())
//
//}

type Test struct {
	base
}

type base struct {
}

func (b *base) Name() string {
	return "base"
}

type Base interface {
	Name() string
}

//func main() {
//	var dict map[string]map[string]int64
//	dict = make(map[string]map[string]int64)
//	addItem(dict)
//	fmt.Println(dict["table"])
//}

func addItem(dc map[string]map[string]int64) {
	if dc == nil {
		fmt.Println("no gc")
	}
	fmt.Println("add item")
	dc["table"] = map[string]int64{"a": 1}
}

func useSortReverseSliceDemo() {
	str := "12"
	for _, item := range strings.Split(str, ",") {
		fmt.Println(item)
	}
	fmt.Println(strings.Split(str, ",")[0])
	fmt.Println(32 << 23)
	s_token := fmt.Sprintf("POST / HTTP/1.0\n\rzs-token: %v", "2rVu6yP6WGZWwF2G1_eLV_ZlIcXtGeNlyx0DNthIXPs")
	fmt.Println(len(s_token))
	s := make([]int, 3, 4)
	a := append(s, 1)
	b := append(s, 2)
	fmt.Println(s, a, b)

	ns, err := net.LookupHost("www.baidu.com")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Err: s%", err.Error())
	}

	for _, n := range ns {
		fmt.Fprint(os.Stdout, "--s%", n, "\n")
	}

	data := `{
		"name": "baidu",
		"url": "http://www.baidu.com",
		"logo": "",
		"related_url": ["www.google.com", "www.sina.com"],
		"user_info": [
			{
				"user_id": 100,
				"user_name": "demo"
			}
		],
		"group_info": null
	}`
	host := new(Host)
	json.Unmarshal([]byte(data), host)
	fmt.Println(host)
}

type Host struct {
	Name       string     `json:"name"`
	Url        string     `json:"url"`
	Logo       string     `json:"logo"`
	RelatedUrl []string   `json:"related_url"`
	UserInfo   []UserInfo `json:"user_info"`
}

type UserInfo struct {
	UserId   int64  `json:"user_id"`
	UserName string `json:"user_name"`
}
