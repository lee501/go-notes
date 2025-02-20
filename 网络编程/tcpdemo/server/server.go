package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"

	"github.com/sirupsen/logrus"
)

func main() {
	listener, err := net.Listen("tcp", ":8090")
	if err != nil {
		logrus.Error("err starting server: ", err)
		os.Exit(1)
	}
	defer listener.Close()

	logrus.Info("server listen on port 8090")

	var wg sync.WaitGroup

	for {
		conn, err := listener.Accept()
		if err != nil {
			logrus.Error("conn accept err: ", err)
			return
		}
		wg.Add(1)
		go handleConnection(conn, &wg)
	}
	wg.Wait()
}

func handleConnection(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	defer conn.Close()

	go func() {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			message := scanner.Text()
			logrus.Info("receive the message from agent", message)
		}
	}()

	writer := bufio.NewWriter(conn)
	for {
		var message string
		fmt.Print("input config to agent: ")
		fmt.Scan(&message)

		if message == "exit" {
			logrus.Info("exit connection...")
			break
		}

		_, err := writer.WriteString(message + "\n")
		if err != nil {
			return
		}
		writer.Flush()
	}
}
