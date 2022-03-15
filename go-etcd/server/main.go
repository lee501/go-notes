package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/sevenelevenlee/go-notes/go-etcd/server/utils"
)

func main() {
	// 1. 创建服务-获取商品详情
	//serviceId := uuid.New().String()      // 服务ID
	serviceName := "productDetailService" // 服务名
	serviceAddress := "192.168.123.178"   // 服务地址
	servicePort := 8081                   // 服务端口

	router := mux.NewRouter()
	server := http.Server{
		Addr:    serviceAddress + ":" + strconv.Itoa(servicePort),
		Handler: router,
	}

	// 模拟商品详情 API
	router.HandleFunc("/product/{id:\\d+}", func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		// todo 获取商品详情
		str := "get product id " + vars["id"]
		writer.Write([]byte(str))
	})

	errChan := make(chan error) // 用于保存错误信息的通道

	// 2. 配置etcd
	service, err := utils.NewService("", serviceName, serviceAddress+":"+strconv.Itoa(servicePort))
	if err != nil {
		errChan <- err
	}

	// 3.将服务注册到etcd
	go func() {
		err := service.RegistryServe()
		if err != nil {
			errChan <- err
		}
	}()

	// 4.监听Http服务 (启动服务)
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			errChan <- err
		}
	}()

	// 监听信号 优雅关闭服务器
	go func() {
		sig := make(chan os.Signal)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-sig)
	}()

	getErr := <-errChan // errChan管道为空时，会一直阻塞
	log.Println("服务器正在关闭...")

	// 关闭Http服务前   需要先回收一些资源，如：redis db 等
	go func() {
		err := service.UnRegister() // 服务反注册 (将服务从etcd中删除,释放资源)
		if err != nil {
			log.Fatal("服务删除失败!", err.Error())
		}
	}()

	// 关闭Http服务
	err = server.Shutdown(context.Background())
	if err != nil {
		log.Fatal("服务关闭失败!", err.Error())
	}
	log.Fatal(getErr) // Fatal 打印输出内容  退出应用程序 defer函数不会执行
}
