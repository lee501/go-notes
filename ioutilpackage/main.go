package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

/*
	func ReadAll(r io.Reader) ([]byte, error)
	func ReadFile(filename string) ([]byte, error)
	func WriteFile(filename string, data []byte, perm os.Filemode) error
	func ReadDir(dirname string) ([]os.FileInfo, error)
*/
//demo
func main() {
	//ioutil.ReadAll
	file, err := os.Open("file")
	defer func() {
		file.Close()
	}()
	if err != nil {
		fmt.Println(err)
		return
	}
	b, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))

	//ioutil.ReadFile
	n, err := ioutil.ReadFile("file")
	fmt.Println(n)
	//ioutil.WriteFile
	ioutil.WriteFile("filename", []byte("111111"), 0666)
	//读取文件夹信息 ioutil.ReadDir()
	fileinfo, err := ioutil.ReadDir("/User")
	for _, file = range fileinfo {
		file.Name()
	}
}
