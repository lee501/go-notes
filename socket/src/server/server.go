package main

import (
	"fmt"
	"net"
	"strings"
)

type User struct {
	Username string
	OtherUsername string
	Msg string
	ServerMsg string
}

var (
	userMap = make(map[string]net.Conn)
	user = new(User)
)
func main() {
	addr, _ := net.ResolveTCPAddr("tcp4", "localhost:8899")
	lis, _ := net.ListenTCP("tcp4", addr)

	for {
		conn, _ := lis.Accept()
		go func() {
			for {
				b := make([]byte, 1024)
				//读取数据
				count, _ := conn.Read(b)
				info := strings.Split(string(b[:count]), "-")

				//保存到用户信息
				user.Username = info[0]
				user.OtherUsername = info[1]
				user.Msg = info[2]
				user.ServerMsg = info[3]
				//将用户存到map中
				userMap[user.Username] = conn
				//接收用户存在
				if re, ok := userMap[user.OtherUsername]; ok && re != nil {
					user.ServerMsg = ""
					n, err := re.Write(b[:count])
					if n == 0 || err!= nil {
						delete(userMap, user.OtherUsername)
						conn.Close()
						re.Close()
						break
					}
				} else {
					user.ServerMsg = "对方不在线"
					conn.Write([]byte(fmt.Sprintf("%s-%s-%s-%s", user.Username, user.OtherUsername, user.Msg, user.ServerMsg)))
				}
			}
		}()
	}
}