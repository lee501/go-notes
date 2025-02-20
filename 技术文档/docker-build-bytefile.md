#### 使用docker运行一个二机制文件
    1. 构建可执行的二机制文件
```cmd
    go build -o bh-go-server-sms server.go
```
    2. 创建dockerfile
```dockerfile
FROM golang:1.12 AS build-env
WORKDIR /ROOT
ADD ./bh-go-server-sms ./bh-go-server-sms
ADD ./config ./config
EXPOSE 80
ENTRYPOINT ["./目标执行文件name"] 
```
##### ENTRYPOINT与CMD的区别
```dockerfile
#使用exec格式的时候：
#docker run传入的参数会覆盖cmd的参数，并且添加加到entrypoint后
ENTRYPOINT ["/bin/echo", "Hello"]
CMD ["World"]
#docker run -it image
#-> hello world
#docker run -it image lee
#-> hello lee 
```
    3.对go二进制文件进行执行权限  
```cgo
chomd 755 bh-go-server-sms

//build镜像
docker build -t imagename .
//build完成查看
docker images
```
    4. 启动容器
```cgo
docker run --name imagesname -d -p 80:80
```
    5.使用docker compose部署, docker-compose up -d
```docker compose
version: "3"
services:
  go-web:
    image: bh-go-server-sms
    ports:
      - "51003:51003"
```