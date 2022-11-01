package main

import (
	"html/template"
	"net/http"
)

func nest(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("view/content.html", "view/header.html", "view/footer.html")
	t.ExecuteTemplate(w, "content", nil)
}
