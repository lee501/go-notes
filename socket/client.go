package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	//创建tcpaddr, 服务端ip和端口
	addr, _ := net.ResolveTCPAddr("tcp4", "localhost:8989")
	//申请链接服务器，协议，本地地址，远程地址
	tcpConn, _ := net.DialTCP("tcp4", nil, addr)
	//向服务端发送数据
	count, _ := tcpConn.Write([]byte("我是你爸爸"))
	fmt.Println("客户端向服务端发送的数据量为:", count)
	//通过休眠测试客户端对象不关闭,服务器是否能接收到对象
	time.Sleep(10 * time.Second)
	//关闭连接
	tcpConn.Close()
	fmt.Println("客户端结束")
	//关闭链接
}
