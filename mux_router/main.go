package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	Name string
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from %s!\n", s.Name)
}

type Router struct {
	*mux.Router
	servers []*Server
}

func (r *Router) AddServer(s *Server) {
	r.servers = append(r.servers, s)
}

func main() {
	//r := &Router{
	//	Router: mux.NewRouter(),
	//}
	//
	//// 创建并添加服务器
	//server1 := &Server{Name: "Server 1"}
	//server2 := &Server{Name: "Server 2"}
	//
	//r.AddServer(server1)
	//r.AddServer(server2)
	//
	//// 设置路由
	//r.Handle("/server2", server2).Methods("GET")
	//r.PathPrefix("/server1").Handler(server1)
	//
	//// 启动 HTTP 服务器
	//fmt.Println("Starting server on port 8080...")
	//http.ListenAndServe(":8080", r)
}
