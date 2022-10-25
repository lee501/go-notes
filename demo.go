package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var req = Req{
		Risk: [][]string{[]string{"risk1", "risk2"}, []string{"reason1", "reason2"}},
	}
	b, _ := json.Marshal(req)
	var resp Resp
	json.Unmarshal(b, &resp)
	if re, ok := resp.Risk.([]interface{}); ok {
		fmt.Printf("-----%#v\n", re[0])
	}
	fmt.Printf("-----%#v\n", 111)

	str := "abcde哈"
	fmt.Println(len([]rune(str)))
	fmt.Println(len(str))
	fmt.Println(len([]byte(str)))

}

type Req struct {
	Risk interface{}
}

type Resp struct {
	Risk interface{}
}
