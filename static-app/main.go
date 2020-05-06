package main

import (
	"fmt"
	"net/http"
)
import "html/template"

func welcome(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("view/index.html")
	t.Execute(w, nil) //第二个参数表示向模版传递的数据
}

func main() {
	server := http.Server{Addr: ":8090"}
	fmt.Println("服务启动")
	/*
	访问url以/static/开头,就会把访问信息映射到指定的目录中
	 */
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", welcome)
	server.ListenAndServe()
}
