package main

import (
	"fmt"
	"os"
)

func main() {
	data := []byte("asdasdasd")
	fl, err := os.OpenFile("/Users/lichunliang/workspace/go/go-notes/text.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)
	if err != nil {
		return
	}
	defer fl.Close()
	_, err = fl.Write(data)
	if err != nil {
		fmt.Println(err)
	}
}