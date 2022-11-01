package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

func show(w http.ResponseWriter, re *http.Request) {
	t, _ := template.ParseFiles("view/index.html")
	t.Execute(w, nil)
}

func download(w http.ResponseWriter, re *http.Request) {
	filename := re.FormValue("filename")
	bytes, err := ioutil.ReadFile("/Users/richctrl/Downloads/" + filename)
	if err != nil {
		w.Write([]byte("文件下载失败，未找到该文件"))
		return
	}
	header := w.Header()
	header.Set("Content-Type", "application/octet-stream")
	header.Set("Content-Disposition", "attachment;filename="+filename)
	w.Write(bytes)
}

func test() {
	server := http.Server{Addr: ":8899"}
	http.HandleFunc("/", index)
	http.HandleFunc("/download", download)
	fmt.Println("服务启动")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
