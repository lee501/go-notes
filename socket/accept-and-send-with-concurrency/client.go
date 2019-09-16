package main

import (
	"fmt"
	"net"
	"strconv"
)

func main() {
	//创建tcp address
	addr, _ := net.ResolveTCPAddr("tcp4", "localhost:8899")
	for i := 0; i < 5; i++ {
		//链接服务器
		tcpconn, _ := net.DialTCP("tcp4", nil, addr)
		//向服务端发送数据
		count, _ := tcpconn.Write([]byte("客户端发送的数据" + strconv.Itoa(i)))
		fmt.Println("客户端向服务端发送的数据量为:", count)
		//接收服务端返回数据
		b := make([]byte, 1024)
		n, _ := tcpconn.Read(b)
		fmt.Println("来自服务端的信息", string(b[:n]))
		//关闭链接
		tcpconn.Close()
	}
}
