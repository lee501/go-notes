package go_file_process

import (
	"fmt"
	"io/ioutil"
)

func ListFiles(path string, level int) {
	fileInfo, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, info := range fileInfo {
		if info.IsDir() {
			for curHier := level; curHier > 0; curHier-- {
				fmt.Printf("|\t")
			}
			fmt.Println(info.Name())
			ListFiles(path+"/"+info.Name(), level+1)
		} else {
			for curHier := level; curHier > 0; curHier-- {
				fmt.Printf("|\t")
			}
			fmt.Println(info.Name())
		}
	}

}
