#### 磁盘分区

* MSDOS(MBR) 分区表格式与限制
    * 开机管理程序记录和分区表都放在第一个扇区（512byte）
        1. MBR Master Boot Record
        2. 分区表 64bytes： 只能做四个Primary分区或Extended分区
           ```
           1. 其实所谓的“分区”只是针对那个64 Bytes的分区表进行设置而已！
           2. 硬盘默认的分区表仅能写入四组分区信息
           3. 这四组分区信息我们称为主要（Primary）或延伸（Extended）分区
           4. 分区的最小单位“通常”为柱面（cylinder）
           5. 当系统要写入磁盘时，一定会参考磁盘分区表，才能针对某个分区进行数据的处理
           ```
           
    
* 开机流程到操作系统之前的动作：
  * BIOS：该硬盘里面去读取第一个扇区的MBR位置；
  * MBR：内含开机管理程序；
  * 开机管理程序（boot loader）：load核心文件来执行的软件；
  * 核心文件：开始操作系统的功能...
    

#### 硬盘挂载
* 查看是否有可挂在硬盘
```shell
    lsblk
        NAME          MAJ:MIN RM  SIZE RO TYPE MOUNTPOINT
        nvme0n1       259:0    0    8G  0 disk
        ├─nvme0n1p1   259:1    0    8G  0 part /
        └─nvme0n1p128 259:2    0    1M  0 part
        nvme1n1       259:3    0 69.9G  0 disk
    df -h
        Filesystem      Size  Used Avail Use% Mounted on
        devtmpfs        7.7G     0  7.7G   0% /dev
        tmpfs           7.7G     0  7.7G   0% /dev/shm
        tmpfs           7.7G  424K  7.7G   1% /run
        tmpfs           7.7G     0  7.7G   0% /sys/fs/cgroup
        /dev/nvme0n1p1  8.0G  4.9G  3.2G  61% /
        tmpfs           1.6G     0  1.6G   0% /run/user/1000
```
* 对空硬盘进行文件格式化
```shell
  sudo mkfs -t xfs nvme1n1
``` 

* 挂载
```shell
  mkdir /data
  mount /dev/nvme1n1 /data
```