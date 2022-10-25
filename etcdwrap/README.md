#### 封装etcd

##### Usage
```go
package main

import (
	"fmt"
	"time"

	"etcdwrap"
)

const (
	etcdHost         = "127.0.0.1:2379"
	registerKey      = "you define etcd key"
	serviceAddr      = "your service addr"
	registerLease    = 5
	renewalLeaseTime = 2 * time.Second
)

func main() {
	//init etcd client
	if _, err := etcdwrap.New(etcdHost); err != nil {
		fmt.Println("etcd init failed")
		return
	}
	//服务注册
	sevice := etcdwrap.ServiceRegister{
		Key:      registerKey,
		Value:    serviceAddr,
		LeaseTTL: registerLease,
		Renewal:  renewalLeaseTime,
	}
	if err := sevice.Start(); err != nil {
		fmt.Println("etcd service register faild")
		return

	}
}

//etcd 监听
func watcher() {
	watcher := wrapetcd.Watcher{
		EtcdKey:  "key",
		Callback: yourCallBack,
	}
	watcher.Register()
}

func yourCallBack(watchkey string, result wrapetcd.WatcherResult) {
	//process your logic
}

//operate
func get() {
	wrapetcd.Get("key")
}

func put() {
	wrapetcd.Put("key", "value")
}
```