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
	level := 2
	for _, child := range fileInfo {
		if level < 0 {
			return s, nil
		}
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
		level--
	}
	return s, nil
}

func main() {
	var s []string
	s, _ = GetAllFile("/Users/lichunliang/workspace/go-db-test", s)
	for _, v := range s {
		fmt.Printf("file slice: %v\n", v)
	}
}