#### 孤儿进程和僵尸进程

    * 孤儿进程(无害)
```text
  父进程退出，子进程还在运行，这些进程称为孤儿进程。
  孤儿进程将被init进程(进程号为1)所收养，由init进程对它们完成状态的收集
```

    * 僵尸进程(有害)
```text
  一个父进程用fork创建子进程， 如果子进程退出， 父进程没有利用wait或waitpid来获取子进程的状态信息，
  那么子进程的状态描述信息(进程号，运行时间，退出状态等)依然保留在系统种，最终导致pid消耗完
```

#### linux查看某个进程的线程
```text
    1. ps -T -p pid （—T开启线程查看）

    2. top 输出某个进程的线程
       top -H -p pid
```

#### linux查看磁盘空间
    * df以磁盘分区为单位查看文件系统
```text
    1. df -h
        Filesystem(文件系统)      Size(分区大小)    Used(已使用容量)    Avail(可用容量)   Use%(已用百分比)    Mounted on(挂载点)    
    2. du -h 指定文件
        显示指定文件所占用空间
```

#### 扫描端口
```text
    nmap -Pn -p1-65535 ip地址
```
    
####查看端口使用
```text
    netstat -nltp
```

####抓包工具
```text
    tcpdump -S -nn -vvv -i lo0 port 8080
    
    -i lo0指定捕获接口localhost
    port 8080 过滤端口8080的流量
    -vvv 打印详细的描述信息
    -S 显示序列号的绝对值
```