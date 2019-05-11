package main

import (
	"io"
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		io.WriteString(w, "<form method=\"POST\" action=\"/upload\" enctype=\"multipart/form-data\">")
	}
}

func main() {
	http.HandleFunc("/upload", Handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
