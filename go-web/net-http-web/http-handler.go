package main

import "net/http"

type FHandler struct{}
type SHandler struct{}

func (f *FHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("第一个handler"))
}

func (s *SHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("第二个handler"))
}

func main() {
	f := FHandler{}
	s := SHandler{}

	server := http.Server{Addr: "localhost:8899"}
	http.Handle("/first", &f)
	http.Handle("/second", &s)
	server.ListenAndServe()
}
