#### tcpdump抓包

##### 基础用法
```shell
tcpdump -i eth0 -nn -s0 -v port 80
```
* -i: 选择网卡
* -nn: 不解析域名和端口，方便直接查看ip和端口
* -s0: 截取全部报文，默认只截取前面96字节
* -v: 显示详细信息
* port 80 抓取80上的流量 
* -e: 显示源和目标mac地址
* -p: 网络接口非混杂模式,此时网卡只接受来自网络端口的目的地址指向自己的数据
* -l 可以将输出立即发送给其他命令
* -c 选项来开启数据包缓冲模式
* -w 选项用来把数据报文输出到文件

##### 抓取特定协议的数据
```shell
tcpdump -i eth0 udp
```
* tcp 与 protocol 6 相同
* udp 或 protocol 17

##### 抓取特定主机的数据
* host 可以抓取特定目的地和源 IP 地址的流量
```shell
tcpdump -i eth0 host 10.10.1.1
```
* 使用 src 或 dst 只抓取源或目的地
```shell
tcpdump -i eth0 dst 10.10.1.20
```
* 组合src host, 表示抓取源地址 host的流量
```shell
tcpdump -i eth0 src host 10.10.1.20
```

##### 抓取的文件写入文件
```shell
tcpdump -i eth0 -s0 -w test.pcap
```
* 切割pcap文件, 每3600创建一个文件,每个文件不超过200*1000000字节
```shell
tcpdump -w /tmp/capture-%H.pcap -G 3600 -C 200
```

##### 理解tcpdump的输出
* tcp报文
```text
[S] : SYN（开始连接）
[.] : 没有 Flag
[P] : PSH（推送数据）
[F] : FIN （结束连接）
[R] : RST（重置连接）
```
```shell
21:27:06.995846 IP (tos 0x0, ttl 64, id 45646, offset 0, flags [DF], proto TCP (6), length 64)
    192.168.1.106.56166 > 124.192.132.54.80: Flags [S], cksum 0xa730 (correct), seq 992042666, win 65535, options [mss 1460,nop,wscale 4,nop,nop,TS val 663433143 ecr 0,sackOK,eol], length 0

21:27:07.030487 IP (tos 0x0, ttl 51, id 0, offset 0, flags [DF], proto TCP (6), length 44)
    124.192.132.54.80 > 192.168.1.106.56166: Flags [S.], cksum 0xedc0 (correct), seq 2147006684, ack 992042667, win 14600, options [mss 1440], length 0

21:27:07.030527 IP (tos 0x0, ttl 64, id 59119, offset 0, flags [DF], proto TCP (6), length 40)
    192.168.1.106.56166 > 124.192.132.54.80: Flags [.], cksum 0x3e72 (correct), ack 2147006685, win 65535, length 0

```
* 源地址/端口->目标地址端口： 192.168.1.106.56166 > 124.192.132.54.80
* 第一行[S]: 开始连接
* 第二行[S.]:SYN-ACK
* 第三行[.] 确认ACK

##### 端口扫描
```shell
tcpdump -nn
```

##### 抓取 SMTP/POP3 协议的邮件
* 只提取电子邮件的收件人
```shell
tcpdump -nn -l port 25 | grep -i 'MAIL FROM\|RCPT TO'
```

##### 提取cookies
* Set-Cookie（服务端的 Cookie）和 Cookie（客户端的 Cookie）
```shell
tcpdump -nn -A -s0 -l | egrep -i 'Set-Cookie|Host:|Cookie:'
```

##### 提取 HTTP POST 请求中的密码
```shell
tcpdump -s 0 -A -n -l | egrep -i "POST /|pwd=|passwd=|password=|Host:"
```