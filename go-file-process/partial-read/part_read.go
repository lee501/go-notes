package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

var path = "/Users/lee/Downloads/tencent.dmg"

func readBlock(filepath string) {
	start := time.Now()
	file, err := os.Open(filepath)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()
	//设置块大小
	buffer := make([]byte, 1024)
	//循环读取，处理读到数据的逻辑
	for {
		n, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			log.Println(err)
		}
		if n == 0 {
			break
		}
		//处理读取到的数据
	}
	fmt.Println("the spend time :", time.Now().Sub(start))
}
