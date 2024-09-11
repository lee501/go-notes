package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"sync"
)

var lock sync.Mutex
var trueList []string
var ip string
var list string

func main() {
	flag.StringVar(&ip, "l", ":8899", "-l=0.0.0.0:8899 指定服务监听的端口")
	flag.StringVar(&list, "d", "182.61.200.7:443", "-d=182.61.200.7:443 指定后端的IP和端口,多个用','隔开")
	flag.Parse()
	trueList = strings.Split(list, ",")
	fmt.Println(trueList)
	if len(trueList) <= 0 {
		fmt.Println("后端IP和端口不能空,或者无效")
		os.Exit(1)
	}
	server()
}
func server() {
	lis, err := net.Listen("tcp", ip)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer lis.Close()
	for {
		conn, err := lis.Accept()
		if err != nil {
			fmt.Println("建立连接错误:%v\n", err)
			continue
		}
		fmt.Println(conn.RemoteAddr(), conn.LocalAddr())
		go handler(conn)
	}
}
func handler(sconn net.Conn) {
	defer sconn.Close()
	ip, ok := getIP()
	if !ok {
		return
	}
	dconn, err := net.Dial("tcp", ip)
	if err != nil {
		fmt.Printf("连接%v失败:%v\n", ip, err)
		return
	}
	fmt.Printf("连接%v成功\n", dconn)
	var wait sync.WaitGroup
	wait.Add(2)
	go func(sconn net.Conn, dconn net.Conn, wait *sync.WaitGroup) {
		_, err := io.Copy(dconn, sconn)
		fmt.Printf("往%v发送数据失败:%v\n", ip, err)
		wait.Done()
	}(sconn, dconn, &wait)
	go func(sconn net.Conn, dconn net.Conn, wait *sync.WaitGroup) {
		_, err := io.Copy(sconn, dconn)
		fmt.Printf("从%v接收数据失败:%v\n", ip, err)
		wait.Done()
	}(sconn, dconn, &wait)
	wait.Wait()
	dconn.Close()
}
func getIP() (string, bool) {
	lock.Lock()
	defer lock.Unlock()
	if len(trueList) < 1 {
		return "", false
	}
	ip := trueList[0]
	trueList = append(trueList[1:], ip)
	return ip, true
}
