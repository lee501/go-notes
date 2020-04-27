package main

import (
	"fmt"
	"net/http"
)

func welcome(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html;charset=usf-8")
	fmt.Fprintln(res, "Welcome to <b>go</b>")
}

func main() {
	http.HandleFunc("/", welcome)
	http.ListenAndServe("localhost:3301", nil)
	fmt.Println("服务已启动")
}

