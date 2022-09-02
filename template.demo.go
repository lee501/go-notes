package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"os"
	"time"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"-""`
}

func main() {
	p := Person{"longshuai", 23}
	tmpl, _ := template.New("test").Parse("Name: {{.Name}}, Age: {{.Age}}")
	_ = tmpl.Execute(os.Stdout, p)

	var m Person
	str := `{"name": "lee", "age": 23}`
	json.Unmarshal([]byte(str), &m)
	fmt.Println(m, m.Age)

	s := "1.xlsx"
	fmt.Println(s[0])

	t := "2022-08-29 23:59:59"
	t1, _ := time.ParseInLocation("2006-01-02 15:04:05", t, time.Local)
	fmt.Println(t1.Unix())
}
