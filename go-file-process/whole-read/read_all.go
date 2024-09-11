package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

var path = "/Users/lee/Downloads/tencent.dmg"

func readAll(filepath string) {
	start := time.Now()
	//打开文件
	file, err := os.Open(path)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()
	//获取文件信息
	fileInfo , err := file.Stat()
	if err != nil {
		log.Println(err)
		return
	}
	//设置buffer大小
	buffer := make([]byte, fileInfo.Size())
	//读取文件到buffer中
	n, err := file.Read(buffer)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(buffer[:n]))
	fmt.Println("the spend time :", time.Now().Sub(start))
}
