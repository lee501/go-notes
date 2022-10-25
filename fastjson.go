package main

import (
	"encoding/json"
	"fmt"

	"github.com/valyala/fastjson"
)

func main() {
	str := []string{"openid123"}
	param := []string{"openid", "0"}
	b, err := json.Marshal(str)
	if err == nil {
		v := fastjson.GetString(b, param[1:]...)
		if v == "" {
			fmt.Println("nil")
		}
		fmt.Println(v)
	}

	str1 := `{"Mobile":"13021015699"}`
	param = []string{"Mobile"}
	var p fastjson.Parser
	v, _ := p.Parse(str1)
	fmt.Printf("mobile=%s\n", v.GetStringBytes(param...))
}
