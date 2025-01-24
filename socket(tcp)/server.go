package main

import (
	"fmt"
	"net"
)

/*
	*TCPAddr结构体包括服务器IP和端口
		type IP []byte
		Port 服务器监听的接口

	type TCPAddr struct {
		IP IP
		Port int
		Zone string
	}
	*TCPConn is an implementation fo Conn interface for TCP network
	tyoe TCPConn struct {
		conn
	}

	type TCPListener struct {
		fd *netFD
	}
*/
func tcpServe() {
	//创建TCPAddr，指定tcp协议，和服务端端口
	addr, _ := net.ResolveTCPAddr("tcp4", "localhost:8989")
	//监听TCPAddr设置的地址
	tcplister, _ := net.ListenTCP("tcp4", addr)

	fmt.Println("服务器已经启动")
	//接收客户端发送的消息，阻塞状态
	tcpConn, _ := tcplister.Accept()
	//读取数据
	b := make([]byte, 1024)
	n, _ := tcpConn.Read(b)
	fmt.Println("接收到的数据：", string(b[:n]))
	//关闭链接
	tcpConn.Close()
}
