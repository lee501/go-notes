package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func GetAllFile(filepath string, s []string) ([]string, error){
	fileInfo, err := ioutil.ReadDir(filepath)
	if err != nil {
		log.Println("read dir failed: ", err)
		return s, nil
	}

	for _, child := range fileInfo {
		if child.IsDir() {
			fullpath := filepath + "/" + child.Name()
			s, err = GetAllFile(fullpath, s)
			if err != nil {
				log.Println("read dir failed: ", err)
				return s, nil
			}
		} else {
			fullName := filepath + "/" + child.Name()
			s = append(s, fullName)
		}
	}
	return s, nil
}

func main() {
	var s []string
	s, _ = GetAllFile("/Users/richctrl/WeChatProjects/ADG/", s)

	fmt.Printf("file slice: %v", s)
}