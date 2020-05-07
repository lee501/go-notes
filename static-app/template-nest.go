package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func nest(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("view/content.html", "view/header.html", "view/footer.html")
	t.ExecuteTemplate(w, "content", nil)
}

func main() {
	server := http.Server{Addr: ":8899"}
	http.HandleFunc("/", nest)
	fmt.Println("8899启动")
	server.ListenAndServe()
}
