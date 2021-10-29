####docker

* 实现一个资源隔离的容器
    1. 
##### namespace
    * Namespaces是linux提供的用于分离进程树、网络接口、挂载点、以及进程间通信等资源的方法。
      docker通过linux的namespaces对不同的容器实现了隔离。

```text
 隔离：   
    进程树
    网络接口
    挂载点
    进程间通信等
```

    * linux命名空间机制提供7种不同命名空间:
      包括CLONE_NEWCGROUP、CLONE_NEWIPC、
      CLONE_NEWNET、CLONE_NEWNS、CLONE_NEWPID、
      CLONE_NEWUSER、CLONE_NEWUTS. 
      创建新的进程时，设置新进程在哪些资源与宿主机隔离。
```text
    名称         宏定义              隔离的资源
    Cgroup      CLONE_NEWCGROUP     cgroup的根目录
    IPC         CLONE_NEWIPC        隔离System IPC(信号量、消息队列和共享内存)和POSIX消息队列
    NET         CLONE_NEWNET        隔离网络资源(网络设备、网络栈、端口等)
    PID         CLONE_NEWPID        隔离进程ID
    Mount       CLONE_NEWNS         隔离Mount Points文件系统挂载点
    User        CLONE_NEWUSER       隔离用户和用户组
    UTS         CLONE_NEWUTS        隔离主机和域名
```

    * 查看当前进程所属的namespace
```text
    ll /proc/$$/ns
```

##### 网络
    
    * 每一个使用docker run启动的容器都具有单独的网络命名空间， docker 提供了四种网络模式

```text
    1. HOST
    2. CONTAINER
    3. NONE
    4. BRIDGE
```
    
    * 默认为网桥模式
        在这种模式下，除了分配隔离的网络命名空间外，docker会为每个容器创建IP地址，
        当docker服务器在主机上启动之后会创建新的虚拟网桥docker0, 并将docker0的ip设置为默认网关
        
        每个容器创建时会创建一对虚拟网卡，组成数据的通道， 一个放在container中，另一个放在docker0
```text
        |container|                     |container|
        |   eth1  |                     |   eth1  | 
             |                               | 
             |                               |
        |- 虚拟网卡pl9867eth1———虚拟网卡pl10744eth1—｜
        |               虚拟网桥docker0             ｜
                               |
                            iptables
                               |
                              eth0

        brctl show查看当前网桥的情况
        
        Docker 通过 Linux 的命名空间实现了网络的隔离，
        又通过 iptables 进行数据包转发，让 Docker 容器能够优雅地为宿主机器或者其他容器提供服务
```   
                  
* Docker 是如何将容器的内部的端口暴露出来并对数据包进行转发的
```text
    当有 Docker 的容器需要将服务暴露给宿主机器，就会为容器分配一个 IP 地址，同时向 iptables 中追加一条新的规则
```

#### 挂载点
```text
    在新的进程中创建隔离的挂载点命名空间，需传入CLONE_NEWNS, 这样子进程就可以获得父进程挂载点的拷贝，
    对文件系统的读写不会同步回父进程以及整个主机的文件系统

    一个容器的启动，一定需要提供一个跟文件系统rootfs， 容器需要使用这个文件系统来创建新的进程，所有二进制执行都必须在
    这个根文件系统中。
    在rootfs中挂载特定的目录
        /proc
        /dev
        /dev/shm
        /dev/mqueue
        /dev/pts
        /sys

     1. mount()   不能访问宿主机的其他目录
     2. chroot    
        在 Linux 系统中，系统默认的目录就都是以 / 也就是根目录开头的，
        chroot 的使用能够改变当前的系统根目录结构，通过改变当前系统的根目录，
        我们能够限制用户的权利，在新的根目录下并不能够访问旧系统根目录的结构个文件，
        也就建立了一个与原系统完全隔离的目录结构
        // chroot
        mount(rootfs, "/", NULL, MS_MOVE, NULL);
        chroot(".");
        chdir("/");
```

####CGroups
```text
    隔离宿主机的物理资源（CPU、MEMORY）
```
1. 每一个CGroup是一组被相同标准和参数限制的进程，不同的cgroup具有层级关系。
   linux的CGroups能够为一组进程分配资源
```text
    Linux 使用文件系统来实现 CGroup，我们可以直接使用下面的命令查看当前的 CGroup 中有哪些子系统
    
    
    $ lssubsys -m 
    cpuset /sys/fs/cgroup/cpuset
    cpu /sys/fs/cgroup/cpu
    cpuacct /sys/fs/cgroup/cpuacct
    memory /sys/fs/cgroup/memory
    devices /sys/fs/cgroup/devices
    freezer /sys/fs/cgroup/freezer
    blkio /sys/fs/cgroup/blkio
    perf_event /sys/fs/cgroup/perf_event
    hugetlb /sys/fs/cgroup/hugetlb
```
* Linux 的命名空间和控制组分别解决了不同资源隔离的问题，前者解决了进程、网络以及文件系统的隔离，后者实现了 CPU、内存等资源的隔离

##### docker 镜像
    Docker 镜像其实本质就是一个压缩包，我们可以使用下面的命令将一个 Docker 镜像中的文件导出：
```text
    $ docker export $(docker create busybox) | tar -C rootfs -xvf -
    $ ls
    bin  dev  etc  home proc root sys  tmp  usr  var
```

##### 容器和镜像的区别
```text
1. Docker 中的每一个镜像都是由一系列只读的层组成的，Dockerfile 中的每一个命令都会在已有的只读层上创建一个新的层
2. 当镜像被 docker run 命令创建时就会在镜像的最上层添加一个可写的层，也就是容器层，
   所有对于运行时容器的修改其实都是对这个容器读写层的修改。

容器和镜像的区别就在于，所有的镜像都是只读的，而每一个容器其实等于镜像加上一个可读写的层，也就是同一个镜像可以对应多个容器
```
