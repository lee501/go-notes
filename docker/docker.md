# docker镜像
```gotemplate
#源镜像
FROM golang:latest
#作者
MAINTAINER lee
#设置工作目录 mydocker为容器名称
WORKDIR /workspace/mydocker
#当前目录下的文件添加到docker容器中
ADD . /workspace/mydocker
#go构建可执行文件
RUN go build .
#暴露端口
EXPOSE 8091
#最终运行docker的命令
ENTRYPOINT  ["./mydocker"]

```
# docker命令
```gotemplate
构建镜像
docker build -t name .
镜像tag
docker tag name newname

运行docker
docker run -p 80:80 -d name

查看镜像
docker images

查看容器
docker ps

根据容器id查看容器的信息
docker inspect id

```