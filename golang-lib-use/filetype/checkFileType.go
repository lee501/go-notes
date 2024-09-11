package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/h2non/filetype"
)

func main() {
	buf, _ := ioutil.ReadFile("/Users/lee/workspace/pcapfiles/files/test.docx")

	kind, _ := filetype.Match(buf)
	if kind == filetype.Unknown {
		fmt.Println("Unknown file type")
		return
	}

	fmt.Printf("File type: %s. MIME: %s\n", kind.Extension, kind.MIME.Value)

	test := "xlsx,pdf,txt,docx,pptx,zip,rar,7z,odt,rtf"
	if strings.Contains(test, "") {
		fmt.Println("1")
	}
	type UserRequest struct {
		BodyMap string `json:"Body"`
	}

	p := UserRequest{
		BodyMap: "我是body map",
	}
	obj, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(obj))

}
