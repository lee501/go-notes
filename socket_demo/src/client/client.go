package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
)

type User struct {
	Username      string
	OtherUsername string
	Msg           string
	ServerMsg     string
}

var (
	user = new(User)
	wg   sync.WaitGroup
)

func main() {
	wg.Add(1)
	fmt.Println("请登陆, 输入用户名: ")
	fmt.Scanln(&user.Username)
	fmt.Println("请输入要给谁发消息：")
	fmt.Scanln(&user.OtherUsername)

	//创建tcp addr
	addr, _ := net.ResolveTCPAddr("tcp4", "localhost:8899")
	//创建conn
	conn, _ := net.DialTCP("tcp4", nil, addr)
	//发送消息
	go func() {
		fmt.Println("请输入: (只启动时提示一次)")
		//循环处理
		for {
			fmt.Scanln(&user.Msg)
			if user.Msg == "exit" {
				conn.Close()
				wg.Done()
				os.Exit(0)
			}
			_, _ = conn.Write([]byte(fmt.Sprintf("%s-%s-%s-%s", user.Username, user.OtherUsername, user.Msg, user.ServerMsg)))
		}
	}()
	//接收消息
	go func() {
		for {
			b := make([]byte, 1024)
			n, _ := conn.Read(b)
			user2 := new(User)
			arrStr := strings.Split(string(b[:n]), "-")
			user2.Username = arrStr[0]
			user2.OtherUsername = arrStr[1]
			user2.Msg = arrStr[2]
			user2.ServerMsg = arrStr[3]

			if user2.ServerMsg != "" {
				fmt.Println("\t\t\t服务器消息:", user2.ServerMsg)
			} else {
				fmt.Println("\t\t\t", user2.Username, ":", user2.Msg)
			}
		}
	}()
	wg.Wait()
}
