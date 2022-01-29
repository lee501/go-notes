package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

type P struct {
	Name string
	Age int
}
func table(wr http.ResponseWriter, re *http.Request) {
	t, _ := template.ParseFiles("view/table.html")
	t.Execute(wr, nil)
}

func getUsers(wr http.ResponseWriter, re *http.Request) {
	users := make([]P, 0)
	users = append(users, P{"lee", 20})
	users = append(users, P{"Ying", 19})
	wr.Header().Set("Content-Type", "application/json;charset=utf-8")
	b, _ := json.Marshal(users)
	fmt.Fprintln(wr, string(b))
}

func main() {
	server := http.Server{Addr: ":8899"}
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/table", table)
	http.HandleFunc("/getUsers", getUsers)
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
