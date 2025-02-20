package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// 处理从服务端接收消息的函数
func receiveMessages(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Printf("Received from server: %s\n", scanner.Text())
	}
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8090")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}
	defer conn.Close()

	// 启动一个 goroutine 来接收消息
	go receiveMessages(conn)

	writer := bufio.NewWriter(conn)
	for {
		var message string
		fmt.Print("Send message to server: ")
		fmt.Scanln(&message)

		// 发送消息
		writer.WriteString(message + "\n")
		writer.Flush()
	}
}
