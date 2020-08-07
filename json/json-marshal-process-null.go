package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Roles []Role `json:"roles"`
}

type Role struct {
	Name string
}
func main() {
	//初始化Role切片 避免返回null
	str := ""
	fmt.Println(str)
	user := User{Name: "lee", Roles: []Role{{Name: "test"}, {Name: "test2"}}}
	data, _ := json.Marshal(&user)
	fmt.Println(string(data))
	b := `{"name":"lee","roles":[{"Name":"test"},{"Name":"test2"}]}`
	var u User
	err := json.Unmarshal([]byte(b), &u)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(u)

	l := []int{1,2,3,4,5}
	for _, v := range l {
		if v >3 {
			continue
		}
		fmt.Println(v)
	}
}
