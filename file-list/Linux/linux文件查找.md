#### 文件搜索

* which 查找可执行文件
```shell
 > which [-a] command 
  * a 将所有由 PATH 目录中可以找到的指令均列出, 而不止第一个被找到的指令名称
  * command 执行文件
```
* whereis 查找文件
```shell
 > whereis 文件名
```

* find 
```shell
 #与时间有关的选项
 > find [path] [option] [action]
  示例 find / -mtime 2 #2天之前更改的文件
  #某个人的文件
  > find / -user username
  #查找某个文件
  > find / -name 文件名
  #查找文件大小
  > find / -size +1M
```
