package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func readByLine(filepath string) {
	start := time.Now()
	file, err := os.Open(filepath)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		fmt.Println(string(line))
	}
	fmt.Println("readEachLine spend : ", time.Now().Sub(start))
}

func readLineByScanner(filepath string) {
	start := time.Now()
	fileHandle, err := os.Open(filepath)
	if err != nil {
		log.Println(err)
		return
	}
	defer fileHandle.Close()
	lineScanner := bufio.NewScanner(fileHandle)
	for lineScanner.Scan() {
		// 如下代码打印每次读取的文件行内容
		fmt.Println(lineScanner.Text())
	}
	fmt.Println("readLineByScanner spend : ", time.Now())
}