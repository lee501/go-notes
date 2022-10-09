package main

import (
	"bufio"
	"io"
	"net"

	"github.com/sevenelevenlee/go-notes/tcp粘包处理/proto"

	log "github.com/sirupsen/logrus"
)

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:10001")
	if err != nil {
		log.Info("listen failed: ", err)
		return
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Info("accept failed: ", err)
			continue
		}
		go processConn(conn)
	}
}

func processConn(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		msg, err := proto.Decode(reader)
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Info("decode msg failed, err:", err)
			return
		}
		log.Info("收到client发来的数据：", msg)
	}
}
