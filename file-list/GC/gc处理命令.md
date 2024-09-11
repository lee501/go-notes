go build --gcflags=-m  逃逸分析
//pprof
go tool pprof -http=:8000 http://localhost:8080/debug/pprof/heap    查看内存使用
go tool pprof -http=:8000 http://localhost:8080/debug/pprof/profile 查看cpu占用

go build -o main
GODEBUG=gctrace=1  打印gc信息

方式2：go tool trace
在代码中添加：
f, _ := os.Create("trace.out")
defer f.Close()
trace.Start(f)
defer trace.Stop()
然后运行程序，等程序运行完成，执行以下命令
go tool trace trace.out