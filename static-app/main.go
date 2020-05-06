package main

import (
	"fmt"
	"net/http"
)
import "html/template"

type User struct {
	Name string
	Age  int
}
func welcome(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("view/index.html")
	//t.Execute(w, "lee") //第二个参数表示向模版传递的数据
	//传递结构体
	//t.Execute(w, User{"lee", 20} )
	//传递map
	m := make(map[string]interface{})
	m["user"] = User{"anne", 21}
	m["money"] = 100
	t.Execute(w, m)
}

func main() {
	server := http.Server{Addr: ":8090"}
	fmt.Println("服务启动")
	/*
	访问url以/static/开头,就会把访问信息映射到指定的目录中
	 */
	//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", welcome)
	server.ListenAndServe()
}
