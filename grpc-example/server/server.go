package main

import (
	"context"
	"fmt"
	"github.com/seveneleven/go-notes/grpc-example/hello"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	address = ":8848"
)

type Server struct {

}

func (s *Server)SayHello(ctx context.Context,in *hello.HelloRequest)(*hello.HelloReply,error){
	return &hello.HelloReply{
		Message:"hello," + in.Name,
	}, nil
}
func main() {
	conn, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("grpc server listening at: 8848 port")
	server := grpc.NewServer()
	hello.RegisterHelloServer(server, &Server{})
	server.Serve(conn)
}
