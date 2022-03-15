package main

import (
	"context"
	"log"

	"github.com/sevenelevenlee/go-notes/go-etcd/client/transport"

	"github.com/sevenelevenlee/go-notes/go-etcd/client/utils"
)

func main() {
	client, err := utils.NewClient()
	if err != nil {
		// etcd 配置出错
		log.Println("etcd连接错误:", err.Error())
	}

	err = client.LoadService() // 加载/services/域下的所有服务
	if err != nil {
		// 没有服务
		log.Println("没有找到服务")
	}
	// 使用随机算法获取服务
	serviceInfo := client.GetService("productDetailService")

	// 调度方法 serviceInfo 服务信息 | "GET" 请求方法为get | services.ProductDetailEncode 商品详情的请求解析方法
	endpoint := client.Call(serviceInfo, "GET", transport.ProductDetailEncode)

	// 执行调度方法         192.168.123.178:8081/product/100
	resp, err := endpoint(context.Background(), transport.ProductDetailRequest{ProductId: 100})

	if err != nil {
		log.Fatal("请求失败:", err.Error())
	}

	log.Println(resp)
}
