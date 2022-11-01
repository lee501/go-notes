package main

import (
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, re *http.Request) {
	t, _ := template.ParseFiles("view/cookie.html")
	t.Execute(w, nil)
}

func getCookie(w http.ResponseWriter, re *http.Request) {
	//c, _ := re.Cookie("name")
	c := re.Cookies()
	t, _ := template.ParseFiles("view/cookie.html")
	t.Execute(w, c)
}

func setCookie(w http.ResponseWriter, re *http.Request) {
	c := http.Cookie{Name: "name", Value: "lee"}
	http.SetCookie(w, &c)
	t, _ := template.ParseFiles("view/cookie.html")
	t.Execute(w, nil)
}
