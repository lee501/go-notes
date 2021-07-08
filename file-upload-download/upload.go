package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
)

func index(w http.ResponseWriter, request *http.Request) {
	t, _ := template.ParseFiles("view/index.html")
	t.Execute(w, nil)
}

func upload(w http.ResponseWriter, request *http.Request) {
	filename := request.FormValue("filename")
	file, fileHeader, _ := request.FormFile("file")
	n := strings.LastIndex(fileHeader.Filename, ".")
	suffix := fileHeader.Filename[n:]
	buffer, _ := ioutil.ReadAll(file)
	ioutil.WriteFile("/Users/lee/Workspace/"+filename+suffix, buffer, 0777)
	t, _ := template.ParseFiles("view/success.html")
	t.Execute(w, filename)
}

func main() {
	server := http.Server{Addr: ":8899"}
	http.HandleFunc("/", index)
	http.HandleFunc("/upload", upload)
	fmt.Println("服务启动")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
