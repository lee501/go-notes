package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"sync"
)

//func main() { // 没有直接在 main 写是因为把统一的操作封装在一个函数中比较利于以后的扩展
//	forword() // 转发函数
//}

func forword() {
	lis, err := net.Listen("tcp", "127.0.0.1:3389") // 本地监听的端口
	if err != nil {
		log.Fatalln("端口监听失败 -> ", err) // 因为端口监听失败所以意味着程序无法使用，所以直接退出程序  log.Fatalln = log.Println + os.Exit 因为监 听未成功所以也不需要 Close()
	}
	defer lis.Close() // 这个函数可能永远都 不会执行，不过还是写上比较好

	for {
		localConn, err := lis.Accept() // 开始接受连接
		fmt.Println(1, localConn.RemoteAddr(), localConn.RemoteAddr())
		if err != nil {
			log.Println(err)
			continue // 部分连接出错不 会影响使用性所以继续执行
		}
		go handle(localConn) // 开始转发， 为了各个链接互不干扰所以使用 go 关键字 新建线程进行处理
	}
}

func handle(localConn net.Conn) {
	var wg sync.WaitGroup
	remoteConn, err := net.Dial("tcp", "182.61.200.7:443") // 转发到的 ip 地址，以及端口，请替换为你需要和目标地址
	fmt.Println(2, remoteConn.RemoteAddr())
	if err != nil {
		localConn.Close()            // 远程 地址链接失败所以，本地监听也没有意义，所以直接关闭 掉
		log.Fatalln("远程链接建立失败", err) // 打印错误并退出程序
	}

	wg.Add(2)
	go func(local net.Conn, remote net.Conn) {
		defer wg.Done()
		io.Copy(remote, local) // 转发数据
		fmt.Println("3")
		remote.Close() // 关闭连接 防止浪费
	}(localConn, remoteConn)
	go func(local net.Conn, remote net.Conn) {
		defer wg.Done()
		io.Copy(local, remote) // 转发数据
		fmt.Println("4")
		local.Close() // 关闭连接 防止浪费
	}(localConn, remoteConn)
	fmt.Println("finish")
	wg.Wait() // 等待数据转发的完成
}
