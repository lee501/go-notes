package main

import (
	"fmt"
	"net"

	log "github.com/sirupsen/logrus"

	"github.com/sevenelevenlee/go-notes/tcp_package/proto"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:10001")
	if err != nil {
		log.Info("dial failed, err", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := `Hello, Hello. How are you?`
		data, err := proto.Encode(msg)
		if err != nil {
			fmt.Println("encode msg failed, err:", err)
			return
		}
		conn.Write(data)
	}
}
