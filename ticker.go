package main

import (
	"fmt"
	"time"
)

type TidConfig struct {
	Version           string            `json:"Version"`
	JsVersion         int               `json:"JsVersion"`
	ChunkMode         string            `json:"ChunkMode"`
	ControlFlagStruct map[string]string `json:"ControlFlag"`
}

func main() {
	ticker := time.NewTicker(3 * time.Second)
	i := 0
	for {
		i += 3
		select {
		case <-ticker.C:
			fmt.Println("---------", i)
		}
		if i > 10 {
			break
		}
	}
	ticker.Stop()
	fmt.Println("ticker finish")

	timer := time.NewTimer(2 * time.Second)
	defer timer.Stop()

	if !timer.Stop() {
		<-timer.C
	}
	timer.Reset(2 * time.Second)
	for {
		i += 3
		select {
		case <-timer.C:
			timer.Reset(2 * time.Second)
			fmt.Println("---------", i)
		}
	}

	//js := "{\"Version\":\"o6\",\"JsVersion\":5,\"ChunkMode\":\"UNCHUNKING\",\"ControlFlag\":{\"dynamic\":\"ON\",\"mks\":\"OFF\",\"timing\":\"OFF\"}}"
	//var t TidConfig
	//err := json.Unmarshal([]byte(js), &t)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(t)
}
