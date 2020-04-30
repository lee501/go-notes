package main

import (
	"fmt"
	"net/http"
	"reflect"
)

func main() {
	server := http.Server{Addr: "localhost:8899"}
	http.HandleFunc("/param", param)
	server.ListenAndServe()
}

func param(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("获取header信息"))
	//获取header信息 map[key]sring
	header := r.Header
	fmt.Println(header)
	fmt.Fprintln(w, "header信息： ", header)
	fmt.Println(reflect.TypeOf(header["Accept-Encoding"]).String())
	for _, v := range header["Accept-Encoding"] {
		fmt.Println(v)
	}
	//获取请求参数, 需要先解析
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Fprintln(w, r.Form["test"][0])

	//通过formvalue取值
	fmt.Fprintln(w, r.FormValue("test"))
}
