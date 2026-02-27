# Go Notes

Go 学习笔记与示例代码整理。

## 目录

- [Go 基础](#go-基础)
- [并发编程](#并发编程)
- [网络编程](#网络编程)
- [文件与 IO](#文件与-io)
- [数据处理](#数据处理)
- [数据结构与算法](#数据结构与算法)
- [面向对象与设计模式](#面向对象与设计模式)
- [数据库](#数据库)
- [性能分析](#性能分析)
- [工具库使用](#工具库使用)
- [技术文档](#技术文档)
- [DevOps](#devops)
- [LLM](#llm)
- [其他示例](#其他示例)

---

## Go 基础

| 目录/文件 | 说明 |
|---|---|
| [golang-learn](./golang-learn) | Go 基础语法学习：枚举、切片、Map、Set、链表、defer、时间函数等 |
| [go基础问题](./go基础问题) | Go 常见基础问题整理（20+ 个专题） |
| [defer](./defer) | defer 语句用法示例 |
| [slice](./slice) | 切片操作与底层原理 |
| [struct](./struct) | 结构体详解、匿名组合、unsafe pointer |
| [emum](./emum) | 枚举（iota）用法 |
| [context](./context) | context 包用法与原理 |
| [string-byte-rune](./string-byte-rune) | string / byte / rune 类型转换 |
| [func-Params](./func-Params) | 函数参数传递 |
| [reflect-and-interface](./reflect-and-interface) | 反射与接口 |
| [位操作](./位操作) | 位运算操作 |
| [Fixnum_bit_range](./Fixnum_bit_range) | 定点数位域范围 |
| [字符串零拷贝](./字符串零拷贝) | 字符串零拷贝技巧 |
| [break_demo.go](./break_demo.go) | break/continue 用法 |
| [generate-stack.go](./generate-stack.go) | 栈生成示例 |
| [panic-and-recover](./panic-and-recover) | panic 与 recover 处理 |
| [select-sum](./select-sum) | select 语句示例 |

## 并发编程

| 目录/文件 | 说明 |
|---|---|
| [go并发操作](./go并发操作) | Go 并发操作综合示例 |
| [channel](./channel) | channel 用法、panic/deadlock、select |
| [go-for-select](./go-for-select) | for + select 模式 |
| [Mutex](./Mutex) | 互斥锁使用 |
| [wait_group](./wait_group) | WaitGroup 用法 |
| [share-memory-and-chan](./share-memory-and-chan) | 共享内存 vs channel 通信 |
| [map-concurrent-process](./map-concurrent-process) | Map 并发读写 |
| [producer-and-consumer.go](./producer-and-consumer.go) | 生产者-消费者模式 |
| [sync_wait.go](./sync_wait.go) | sync.WaitGroup 示例 |
| [syncmap.go](./syncmap.go) | sync.Map 示例 |
| [ticker.go](./ticker.go) | Ticker 定时器示例 |
| [golang-lib-use/sync.Atomic.go](./golang-lib-use/sync.Atomic.go) | sync/atomic 原子操作 |
| [golang-lib-use/sync.Map源码.go](./golang-lib-use/sync.Map源码.go) | sync.Map 源码分析 |
| [golang-lib-use/sync.Once.go](./golang-lib-use/sync.Once.go) | sync.Once 用法 |
| [golang-lib-use/sync.Pool.go](./golang-lib-use/sync.Pool.go) | sync.Pool 对象池 |

## 网络编程

| 目录/文件 | 说明 |
|---|---|
| [网络编程](./网络编程) | TCP socket、TCP 粘包处理 |
| [go-web](./go-web) | Gin Web 框架示例 |
| [http表单](./http表单) | HTTP 表单处理 |
| [mux_router](./mux_router) | Gorilla Mux 路由 |
| [websocket](./websocket) | WebSocket 示例 |
| [grpc-example](./grpc-example) | gRPC 示例 |
| [go-socket.io-example](./go-socket.io-example) | Socket.IO 示例 |
| [udp-server](./udp-server) | UDP 服务器 |
| [proxy](./proxy) | 代理服务器 |
| [流式代理整理](./流式代理整理) | HTTP/SSE/WebSocket 流式代理 |
| [HttpResetResponseBody](./HttpResetResponseBody) | HTTP 重置响应体 |
| [gin_middleware_modify_response.go](./gin_middleware_modify_response.go) | Gin 中间件修改响应 |
| [rewrite-res-body-under-gin-middleware](./rewrite-res-body-under-gin-middleware) | Gin 中间件重写响应体 |
| [ip-ping](./ip-ping) | IP Ping 工具 |
| [nginx-weighted-round-robin](./nginx-weighted-round-robin) | Nginx 加权轮询算法 |
| [static-app](./static-app) | 静态文件服务 |
| [backend_demo.go](./backend_demo.go) | 后端服务示例 |

## 文件与 IO

| 目录/文件 | 说明 |
|---|---|
| [文件处理](./文件处理) | 文件读写、多种读取方式、gzip 压缩 |
| [IO.reader](./IO.reader) | io.Reader 用法 |
| [IO.writer](./IO.writer) | io.Writer 用法 |
| [ioutilpackage](./ioutilpackage) | ioutil 包用法 |
| [io-cache](./io-cache) | IO 缓存处理 |
| [file-upload-download](./file-upload-download) | 文件上传下载 |
| [file-watcher](./file-watcher) | 文件监听 |
| [file_chunk_upload_service](./file_chunk_upload_service) | 分块上传服务 |
| [断点续传](./断点续传) | 断点续传实现 |
| [inputReader-read-data-from-control](./inputReader-read-data-from-control) | 从控制台读取数据 |
| [bigfile.go](./bigfile.go) | 大文件处理 |
| [toml-file](./toml-file) | TOML 文件解析 |
| [utf82gbk](./utf82gbk) | UTF-8 与 GBK 编码转换 |
| [Os.signal](./Os.signal) | os/signal 系统信号处理 |
| [Os.user](./Os.user) | os/user 用户信息 |

## 数据处理

| 目录/文件 | 说明 |
|---|---|
| [json](./json) | JSON 编解码 |
| [jsonbenchmark](./jsonbenchmark) | JSON 库性能对比 |
| [json2string.go](./json2string.go) | JSON 转字符串 |
| [test_omitempty_of_json.go](./test_omitempty_of_json.go) | JSON omitempty 测试 |
| [fastjson.go](./fastjson.go) | fastjson 库使用 |
| [sonic.demo.go](./sonic.demo.go) | sonic JSON 库使用 |
| [xmlDemo](./xmlDemo) | XML 解析 |
| [excel.demo](./excel.demo) | Excel 文件处理 |
| [excel_gen.go](./excel_gen.go) | Excel 生成 |
| [pdf](./pdf) | PDF 处理 |
| [pdf2word](./pdf2word) | PDF 转 Word |
| [ORC文件创建](./ORC文件创建) | ORC 文件创建 |
| [image-process](./image-process) | 图片处理 |
| [decodede.go](./decodede.go) | 解码示例 |
| [han.go](./han.go) | 汉字处理 |
| [regexp](./regexp) | 正则表达式 |
| [ahocorasick_demo.go](./ahocorasick_demo.go) | Aho-Corasick 多模式字符串匹配 |
| [template_demo.go](./template_demo.go) | Go template 模板 |
| [captche.go](./captche.go) | 验证码生成 |

## 数据结构与算法

| 目录/文件 | 说明 |
|---|---|
| [数据结构和算法](./数据结构和算法) | 哈希表、LRU、LeetCode、排序算法 |
| [list](./list) | 链表 |
| [ring](./ring) | 环形链表 |
| [golang-sort](./golang-sort) | 排序算法 |
| [map-value-is-struct](./map-value-is-struct) | Map value 为 struct 的注意点 |
| [use_sort_reverse_slice.go](./use_sort_reverse_slice.go) | 切片反向排序 |

## 面向对象与设计模式

| 目录/文件 | 说明 |
|---|---|
| [golang-oop-idea](./golang-oop-idea) | Go 面向对象思想总结 |
| [inherit](./inherit) | 结构体继承 |
| [interface-and-struct-method](./interface-and-struct-method) | 接口与结构体方法 |
| [interface_inherit](./interface_inherit) | 接口继承 |
| [polymorphic](./polymorphic) | 多态实现 |
| [接口解耦](./接口解耦) | 接口解耦设计 |
| [func_option_design](./func_option_design) | 函数式选项设计模式 |
| [golang-lib-use/go-option-design](./golang-lib-use/go-option-design) | Option 设计模式 |
| [go-close-func](./go-close-func) | 闭包函数 |

## 数据库

| 目录/文件 | 说明 |
|---|---|
| [mysql](./mysql) | MySQL 操作 |
| [go-etcd](./go-etcd) | etcd 客户端使用 |
| [etcdwrap](./etcdwrap) | etcd 封装 |
| [go-local-cache](./go-local-cache) | 本地缓存（go-cache） |

## 性能分析

| 目录/文件 | 说明 |
|---|---|
| [golang-gc](./golang-gc) | Go GC 分析 |
| [pprof](./pprof) | pprof 性能分析工具 |
| [pprof.analysis](./pprof.analysis) | pprof 分析结果 |
| [cpu.pprof](./cpu.pprof) | CPU profile 文件 |

## 工具库使用

| 目录/文件 | 说明 |
|---|---|
| [golang-lib-use](./golang-lib-use) | 常用库使用：gin-jwt、filetype、uber/zap、flag 等 |
| [cgo](./cgo) | CGO 使用示例 |
| [go-cmd](./go-cmd) | os/exec 执行命令 |
| [cron.go](./cron.go) | 定时任务（robfig/cron） |
| [viper.go](./viper.go) | Viper 配置文件读取 |
| [go_logrus](./go_logrus) | logrus 日志库 |
| [logger](./logger) | 日志封装 |
| [tequila](./tequila) | Tequila 框架示例 |
| [tika-demo](./tika-demo) | Apache Tika 文件解析 |
| [gimli](./gimli) | Gimli 库示例 |
| [ua.go](./ua.go) | User-Agent 解析 |
| [config.yaml](./config.yaml) | 配置文件示例 |

## 技术文档

| 目录/文件 | 说明 |
|---|---|
| [技术文档](./技术文档) | Docker、Nginx、Redis、MySQL、Kafka、Zookeeper、GC、HTTP、Linux 等技术文档 |
| [doc](./doc) | 文档示例 |
| [docconve](./docconve) | 文档转换 |

## DevOps

| 目录/文件 | 说明 |
|---|---|
| [docker](./docker) | Docker 相关示例 |

## LLM

| 目录/文件 | 说明 |
|---|---|
| [LLM](./LLM) | 大语言模型相关代码示例 |

## 其他示例

| 目录/文件 | 说明 |
|---|---|
| [practice](./practice) | 综合练习 |
| [demo](./demo) | 综合示例 |
| [program_demo](./program_demo) | 程序示例 |
| [work-test](./work-test) | 工作测试用例 |
| [golang-learn/switch-and-select](./golang-learn/switch-and-select) | switch 和 select 综合示例 |
