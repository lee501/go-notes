#### linux目录

##### /root目录，与开机系统有关
  - /bin  :放置的是在单人维护模式下还能够被操作的指令,cat, chmod, chown, date, mv, mkdir, cp, bash等
  - /boot :开机使用的文件
  - /dev  :设备都是以文件的型态存在于这个目录
  - /etc  :系统主要的配置文件存放目录
  - /lib  :lib放置系统开机时会用到的函数库
  - /opt  :第三方软件放置的目录
  - /sbin :开机过程需要的指令， 而本机自行安装的软件产生的系统可执行文件，放置在/usr/local/sbin
  - /proc : virtual filesystem, 数据都是在内存当中， 例如系统核心、行程信息（process）、周边设备的状态及网络状态.
    - /proc/cpuinfo, /proc/dma, /proc/interrupts, /proc/ioports, /proc/net/*
  - /sys : virtual filesystem, 记录核心与系统硬件信息较相关的信息。 包括目前已载入的核心模块与核心侦测到的硬件设备信息

##### /usr (unix software resource): Unix操作系统软件资源所放置的目录
  
  - /usr/bin 所有一般用户使用的指令存放的位置, 比如cd top
  - /usr/lib/ 系统使用的函数库
  - /usr/sbin
  - /usr/local/ 系统管理员在本机自行安装自己下载的软件（非distribution默认提供者），建议安装到此目录， 这样会比较便于管理
  - /usr/libexec/ 存放不被一般使用者惯用的可执行文件或脚本
  - /usr/src/ 源代码建议放置到这里

##### /var(variable): 系统运行相关，主要针对常态性变动的文件，包括高速缓存（cache）、登录文件（log file）以及某些软件运行所产生的文件

  - /var/log/   日志
  - /var/lib/	程序本身执行的过程中，需要使用到的数据文件放置的目录
  - /var/run/	某些程序或者是服务启动后，会将他们的PID放置在这个目录下

##### 原本应该要在根目录 （/） 里面的目录，将他内部数据全部挪到 /usr 里面去，然后进行链接设置
* /bin --> /usr/bin
* /sbin --> /usr/sbin
* /lib --> /usr/lib