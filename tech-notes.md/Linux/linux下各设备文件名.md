###Linux下各设备的文件名

1. 硬件设备-> 所有的硬件设备文件都在/dev这个目录内
```text
    1) 硬盘
    /dev/sd[a-d]: 即有/dev/sda, /dev/sdb, /dev/sdc, /dev/sdd这四个文件
    2) USB 接口
    /dev/usb/lp[0-15] 
    3) 鼠标	
    /dev/input/mouse[0-15] （通用） 
    /dev/mouse （当前鼠标）
    4) 虚拟机
    /dev/vd[a-p] 
```

2. Linux的可执行程序及相关的文件摆放的目录
```text
    /usr
```