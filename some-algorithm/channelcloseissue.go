package main

import (
	"fmt"
)

func main() {
	fmt.Println("running, not deadlock")
	waitQueue := make(chan int)
	waitQueue <- 1
	return

	//server, err := net.Listen("tcp", "127.0.0.1:9001")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//for {
	//	connection, err := server.Accept()
	//	if err != nil {
	//		panic("server")
	//	}
	//	fmt.Printf("Received connection from %s.\n", connection.RemoteAddr())
	//}
}