package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/bytedance/sonic"
)

type TaskConfig struct {
	Name   string    `json:"name"`
	Circle int       `json:"circle"`
	Start  time.Time `json:"start"`
	End    time.Time `json:"end"`
}

func sonicDemo() {
	m := struct {
		Name string
		Sex  string
	}{
		Name: "lee",
		Sex:  "man",
	}
	val, err := sonic.Marshal(&m)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(val))

	//local, _ := time.LoadLocation("Asia/Shanghai")
	//start, _ := time.ParseInLocation("2006-01-02 15:04:05", "2022-08-15 00:00:00", local)
	//end, _ := time.ParseInLocation("2006-01-02 15:04:05", "2022-08-15 23:59:59", local)

	t := TaskConfig{
		Name:   "test",
		Circle: 0,
		Start:  time.Now(),
		End:    time.Now().Add(23 * time.Hour),
	}
	re, _ := json.Marshal(t)
	fmt.Println(string(re))
}
