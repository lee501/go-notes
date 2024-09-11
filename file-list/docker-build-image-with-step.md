#####使用docker构建更小的镜像

1. 选择较小的基础镜像
```text
    FROM golang:alpine
```
2. 分阶段build
```text
    FROM golang:alpine AS build-env
    WORKDIR /app
    ADD . /app
    RUN cd /app && go build -o goapp

    FROM alpine
    RUN apk update && \
       apk add ca-certificates && \
       update-ca-certificates && \
       rm -rf /var/cache/apk/*
    WORKDIR /app
    COPY --from=build-env /app/goapp /app
    EXPOSE 8080
    ENTRYPOINT ./goapp
```
