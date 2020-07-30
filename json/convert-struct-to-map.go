package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type UserInfo struct {
	Name string `json:"name"`
	Age  int	`json:"age"`
}

func main() {
	var data []byte
	u := UserInfo{Name: "lee", Age: 35}
	data, _ = json.Marshal(u)
	fmt.Printf("%#v\n", string(data))
	var m map[string]interface{}
	json.Unmarshal(data, &m)
	fmt.Println(reflect.TypeOf(m["age"]))
	//b, _ := json.Marshal(u)
	//fmt.Println(string(b))
	//var u1 UserInfo
	//json.Unmarshal(b, &u1)
	//
	//var m map[string]interface{}
	//json.Unmarshal(b, &m)
	//fmt.Printf("%#v\n", m)
	//var res interface{}
	//json.Unmarshal(b, &res)
	//fmt.Printf("%#v\n", res)
	//for k, v := range m{
	//	fmt.Printf("key: %#v value: %#v\n", k, v)
	//}

	//ToMap(u, "name")
}

func ToMap(in interface{}, tagName string) (map[string]interface{}, error){
	out := make(map[string]interface{})

	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("ToMap only accept struct or struct point; got %T", v)
	}

	t := v.Type()
	//m := reflect.TypeOf(in)
	for i:=0; i< v.NumField(); i++ {
		fi := t.Field(i)
		fmt.Println(fi.Tag.Get("Name"))
		if tagValue := fi.Tag.Get(tagName); tagValue != "" {
			out[tagValue] = v.Field(i).Interface()
		}
	}
	return out, nil
}
