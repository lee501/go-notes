package main

import "net/http"

func test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("测试handle func 函数"))
}
func handleFunc() {
	server := http.Server{Addr: "localhost:3002"}
	http.HandleFunc("/test", test)
	server.ListenAndServe()
}
