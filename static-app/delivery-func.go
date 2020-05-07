package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

func MyFunc(t string) string {
	m, _ := time.Parse("2006-01-02 15:04:05", t)
	fmt.Println(m)
	m = m.Add(60 * time.Second)
	fmt.Println(m)
	return m.Format("2006-01-02 15:04:05")
}

func funcParam(w http.ResponseWriter, request *http.Request) {
	//把自定义函数绑定到FuncMap上
	funcmap := template.FuncMap{"mfunc": MyFunc}
	t := template.New("test.html").Funcs(funcmap)
	//解析模版
	t, _ = t.ParseFiles("view/test.html")
	time := "2020-05-10 10:50:00"
	t.Execute(w, time)
}

func main() {
	server := http.Server{Addr: ":8080"}
	http.HandleFunc("/", funcParam)
	fmt.Println("服务启动")
	server.ListenAndServe()
}
