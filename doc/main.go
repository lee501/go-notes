package main

import (
	"fmt"
	"os"

	"github.com/lee501/doc"
)

func main() {
	// 打开 Word 文档
	file, err := os.Open("document.doc")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 从文档中提取文本
	text, err := doc.ParseDoc(file)
	if err != nil {
		panic(err)
	}

	// 将读取器转换为字符串并打印
	fmt.Println(text)
}
