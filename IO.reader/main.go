package main

import (
	"fmt"
	"os"
	"strings"
)

//stream流是应用程序和外部资源进行数据交流的纽带
//实现io的Reader接口就属于输入流
/*
	type Reader interface {
		Read(p []byte) (n int, err error)
	}
*/
func main() {
	//字符串流：strings包的NewReader
	/*type Reader struct {
		s        string
		i        int64 // current reading index
		prevRune int   // index of previous rune; or < 0
	}*/
	r := strings.NewReader("Hello Lee")

	b := make([]byte, r.Size())
	n, err := r.Read(b)
	if err != nil {
		fmt.Println("读取失败，", err)
	}
	fmt.Println("读取自己长度", n)
	fmt.Println("流中数据", string(b))

	//文件流：读取文件数据到程序中
	file, err := os.Open("/Users/richctrl/workspace/os/os.txt")
	defer file.Close()
	//读取文件信息
	fileInfo, _ := file.Stat()
	//创建文件大小的字节切片
	b = make([]byte, fileInfo.Size())
	n, _ = file.Read(b)
	fmt.Println("读取自己长度", n)
	fmt.Println("流中数据", string(b))
}
