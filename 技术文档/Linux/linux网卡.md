#### 网卡

1. 查看网卡状态的命令
```shell
# 列出所有网卡接口
ifconfig -a 

# 列出所有网卡的信息, 包括接口名, 状态, 接收和发送的数据包数量, 错误数等
ip -s link

# mac下查看网卡接口信息
netstat -i  
```

2. 查看网卡实时流量
```shell
ifstat -i <interface>
```

3. 查看端口占用
```shell
netstat -tuln 

lsof -i:<port>
```

4. ifconfig列出的网卡参数说明
```text
* <interface>：网卡的接口名，例如 eth0、en0 等。
* <flags>：表示网卡的状态和特性的标志位，常见的标志位包括：
    * UP：网卡已启用并连接。
    * RUNNING：网卡正在运行。
    * BROADCAST：支持广播通信。
    * MULTICAST：支持多播通信。
    * PROMISC：网卡处于混杂模式，可以接收所有经过的数据包。
* <MTU>：最大传输单元，表示网卡支持的最大数据包大小。
* inet <IP>：网卡配置的 IPv4 或 IPv6 地址。
* netmask <subnet>：子网掩码，定义了 IP 地址的网络部分和主机部分。
* broadcast <broadcast>：广播地址，用于向网络中的所有设备广播消息。
* ether <MAC Address>：网卡的物理地址，也称为 MAC 地址。
* <additional information>：可能会显示其他与网卡相关的信息，例如错误计数、接收和发送的数据包数量等。
```
