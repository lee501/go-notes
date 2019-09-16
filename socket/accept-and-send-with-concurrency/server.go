package main

import (
	"fmt"
	"net"
)

func main() {
	//创建tcp address, 指定tcp4协议和8899端口
	address, _ := net.ResolveTCPAddr("tcp4", "localhost:8899")
	//监听地址
	tcplistener, _ := net.ListenTCP("tcp4", address)

	fmt.Println("服务已启动")
	for {

		//等待接收客户端消息
		conn, _ := tcplistener.Accept()
		go func() {
			//读取客户端消息
			b := make([]byte, 1024)
			n, _ := conn.Read(b)
			fmt.Println("服务端接收的数据", string(b[:n]))

			//向客户端发送数据
			conn.Write([]byte("来自服务端的消息"))
			//关闭链接
			conn.Close()
		}()
	}
}
