package main

import (
	"fmt"
	"net"
)

func main() {
	// 要监听的地址和端口
	address := "192.168.1.100:80"

	// 监听UDP数据包
	conn, err := net.ListenPacket("udp", address)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Printf("Listening on %s...\n", address)

	// 创建一个用于读取数据的缓冲区
	buffer := make([]byte, 2048)

	for {
		// 读取数据
		n, addr, err := conn.ReadFrom(buffer)
		if err != nil {
			fmt.Println("Failed to read data:", err)
			continue
		}

		// 输出收到的消息和发送者地址
		fmt.Printf("Received '%s' from %s\n", string(buffer[:n]), addr)

		// 将消息回显给发送者
		_, err = conn.WriteTo(buffer[:n], addr)
		if err != nil {
			fmt.Println("Failed to write data:", err)
		}
	}
}
