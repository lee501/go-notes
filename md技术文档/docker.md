####docker

##### namespace
```text
 隔离：   
    进程树
    网络接口
    挂载点
    进程间通信等资源的方法
```

* linux提供7种命名空间
```text
    CLONE_NEWCGROUP
    _NEWIPC
    _NEWNET
    _NEWPID
    _NEWNS
    _NEWUSER
    _NEWUTS
```
* 网络
```text
    docker 提供了四种网络模式
        1. HOST
        2. CONTAINER
        3. NONE
        4. BRIDGE
```
* 默认为网桥模式
    docker会为每个容器创建IP地址，当docker服务器在主机上启动之后会创建新的虚拟网桥docker0, 并将docker0的ip设置为默认网关
    
    |container|          |container|
    |   eth0  |          |   eth0  |
         |
         |
    ｜- 虚拟网卡——————————————虚拟网卡——｜
    |            docker0              |
                   |
                iptables
                   |
                  eth0
                  
* Docker 是如何将容器的内部的端口暴露出来并对数据包进行转发的
```text
    当有 Docker 的容器需要将服务暴露给宿主机器，就会为容器分配一个 IP 地址，同时向 iptables 中追加一条新的规则
```

#### 挂载点
```text
    创建隔离的挂载点命名空间，需传入CLONE_NEWNS, 这样子进程对文件系统的读写不会同步回父进程以及整个主机的文件系统
     
     1. mount()   不能访问宿主机的其他目录
     2. chroot    chroot 的使用能够改变当前的系统根目录结构，通过chroot改变进程可以访问文件目录的根节点
```

####CGroups
```text
    隔离宿主机的物理资源
```
1. linux的CGroups能够为一组进程分配资源
```text
    lssubsys -m 查看当前CGroup种有哪些子系统
```