package main

import (
	"net/http"
)

func main()  {
	//type DIR string
	//type Integer int
	//type Test struct {
	//	Name string
	//}
	//
	//in := Integer(1)
	//te := DIR("test")
	//test := Test{"lee"}
	//fmt.Println(in)
	//fmt.Println(te)
	//fmt.Println(test)
	h := http.FileServer(http.Dir("."))
	http.ListenAndServeTLS(":8001", "rui.crt", "rui.key", h)
}