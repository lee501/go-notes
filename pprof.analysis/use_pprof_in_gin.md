#### Gin框架中使用pprof

* 在代码中如何添加pprof
```go
package main

import (
	"net/http"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	pprof.Register(app) // 性能

	app.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "test")
	})
	app.Run(":3000")
}

//代码运行之后可以看到系统自动增加了很多/debug/pprof的API
```

/*
   * 客户服务器上采集性能分析
      curl -sK -v http://localhost:8070/debug/pprof/profile?seconds=600 > profile.out
        curl -sK -v http://localhost:8070/debug/pprof/heap?seconds=30 > heap.out
   * 拷贝到本地
      scp root@49.232.10.214:profile11.out /Users/lee
   * 打开web查看
      go tool pprof -http=:8080 profile11.out

*/
* 使用go tool pprof采集数据, 终端窗口输入以下命令
```shell
go tool pprof --seconds 20 http://localhost:3000/debug/pprof/goroutine

或者采用如下写法

go tool pprof http://localhost:3000/debug/pprof/goroutine?second=20
```
##### pprof查看
1. 使用web查看pprof，在命令行最后一行pprof后输入web
    1. 云服务会报没有合适的浏览器
    
* 可以通过采集的文件创建web服务，(Saved profile一行是需要的数据), 在新的终端输入:
```shell
go tool pprof -http://0.0.0.0:3001 /home/ubuntu/pprof/pprof.rumgo.goroutine.001.pb.gz
```

* 需要安装 graphviz
```shell
brew install graphviz # for macos

apt-get install graphviz # for ubuntu

yum install graphviz # for centos
```

2. go tool pprof命令行交互界面
   
    1. top 查看cpu使用情况
    ```text
        flat:当前函数占用CPU的耗时

        flat%:当前函数占用CPU的耗时百分比
        
        sum%:函数占用CPU的累积耗时百分比
        
        cum：当前函数+调用当前函数的占用CPU总耗时
        
        cum%: 当前函数+调用当前函数的占用CPU总耗时百分比
    ```
   
    2. list 函数名命令查看具体的函数分析
    3. pdf命令可以生成可视化的pdf文件