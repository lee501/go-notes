package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

//将文件以string 对象读出
func readFileByString(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		println(err.Error())
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		readString, err := reader.ReadString('\n')
		println(readString)
		if err == io.EOF {
			return
		}
	}
}

//整个文件的内容读到一个字符串里
func readFileIntoString(inputfile, outfile string) {
	input:= inputfile
	out := outfile
	buf, err := ioutil.ReadFile(input)
	if err != nil {
		println(err.Error())
		return
	}

	_ = ioutil.WriteFile(out, buf, 0644)
}

//带缓冲的读取
func readWithBuffer(file *os.File) {
	buf := make([]byte, 1024)

	n, _ := file.Read(buf)
	if n == 0 {
		return
	}
}

//数据是按列排列并用空格分隔的, 按列读取
func readFileWithCol(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		println(err.Error())
		return
	}
	defer file.Close()

	var col1, col2, col3 []string
	for {
		var v1, v2, v3 string
		_, err = fmt.Fscan(file, &v1, &v2, &v3)
		if err != nil {
			break
		}
		col1 = append(col1, v1)
		col2 = append(col2, v2)
		col3 = append(col3, v3)
	}
}

//获取文件名字
func getFileName(path string) string{
	return filepath.Base(path)
}
