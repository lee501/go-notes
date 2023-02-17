package main

import (
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type User struct {
	Name string
	Age  int
}

func welcome(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("view/index.html")
	if err != nil {
		fmt.Println(err)
	}
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
	//usage go run main.go -path="/Users/lee/workspace/gp-notes/go-web/static-app/static"
	var path string
	flag.StringVar(&path, "path", "static/", "static path")
	flag.Parse()

	file, _ := exec.LookPath(os.Args[0])
	pat, _ := filepath.Abs(file)
	ind := strings.LastIndex(pat, string(os.PathSeparator))
	fmt.Println(pat[:ind])
	server := http.Server{Addr: ":8899"}
	/*
		访问url以/static/开头,就会把访问信息映射到指定的目录中
	*/
	//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/wel", welcome)
	//http.HandleFunc("/param", funcParam)
	http.HandleFunc("/nest", nest) //template

	http.HandleFunc("/index", index)
	http.HandleFunc("/setCookie", setCookie)
	http.HandleFunc("/getCookie", getCookie)
	err := server.ListenAndServe()
	fmt.Println("服务启动")
	if err != nil {
		fmt.Println(err)
	}
}

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

func nest(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("view/content.html", "view/header.html", "view/footer.html")
	t.ExecuteTemplate(w, "content", nil)
}
