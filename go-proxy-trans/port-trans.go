package go_proxy_trans

import (
	"io"
	"net"
)

func server() {
	//listen local server 8080
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		return
	}
	for {
		conn, err := listener.Accept()
		if err == nil {
			go handleClientRequest(conn)
		}
	}
}

func handleClientRequest(conn net.Conn) {
	defer conn.Close()
	//connect remote address
	remote_address := "192.168.210.194:80"
	remote, err := net.Dial("tcp", remote_address)
	if err != nil {
		return
	}
	defer remote.Close()
	//双向copy remote和conn
	go io.Copy(remote, conn)
	io.Copy(conn, remote)
}
