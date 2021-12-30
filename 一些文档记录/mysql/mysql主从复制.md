# mysql主从配置

###  master my.cnf 配置文件设置: 172.17.0.3
```my.cnf
    #mysql master1 config 
    [mysqld]
    server-id = 1        # 节点ID，确保唯一
    
    # log config
    log-bin = mysql-bin     #开启mysql的binlog日志功能
    sync_binlog = 1         #控制数据库的binlog刷到磁盘上去 , 0 不控制，性能最好，1每次事物提交都会刷到日志文件中，性能最差，最安全
    binlog_format = mixed   #binlog日志格式，mysql默认采用statement，建议使用mixed
    expire_logs_days = 7                           #binlog过期清理时间
    max_binlog_size = 100m                    #binlog每个日志文件大小
    binlog_cache_size = 4m                        #binlog缓存大小
    max_binlog_cache_size= 512m              #最大binlog缓存大
    binlog-ignore-db=mysql #不生成日志文件的数据库，多个忽略数据库可以用逗号拼接，或者 复制这句话，写多行
    
    auto-increment-offset = 1     # 自增值的偏移量
    auto-increment-increment = 1  # 自增值的自增量
    slave-skip-errors = all #跳过从库错误
    #binlog-do-db指定需要复制的数据库
```
### slave mysql.cnf 配置: 172.17.0.2
```my.cnf
    [mysqld]
    server-id = 2
    log-bin=mysql-bin #生成的二进制日志文件
    #replicate-do-db要同步的数据库
    relay-log = mysql-relay-bin
    replicate-wild-ignore-table=mysql.%
    replicate-wild-ignore-table=test.%
    replicate-wild-ignore-table=information_schema.%
    slave_net_timeout = 10 //设置网络超时时间，即多长时间测试一下主从是否连接
    log-slave-updates：控制 slave 上的更新是否写入二进制日志，默认为0；若 slave 只作为从服务器，则不必启用；若 slave 作为其他服务器的 master，则需启用，启用时需和 log-bin、binlog-format 一起使用，这样 slave 从主库读取日志并重做，然后记录到自己的二进制日志中；


```
## master和slave数据库操作
1.进入master数据库创建，创建复制的用户   
```bash
    CREATE USER 'repl_user'@'%' IDENTIFIED BY 'repl_passwd';
    #语句中的%代表所有服务器都可以使用这个用户，如果想指定特定的ip，将%改成ip即可
```
2.赋予该用户复制的权利
```bash
    grant replication slave on *.* to 'repl_user'@'172.17.0.2'  identified by 'repl_passwd';
    
    FLUSH PRIVILEGES;
```
3.查看master的状态(//获取file 和 position)
```bash
    flush tables with read lock; 此处需要锁定， master不得进行数据操作(新装机不需要)
    show master status;
    mysql> show master status;
    +------------------+----------+--------------+------------------+-------------------+
    | File             | Position | Binlog_Do_DB | Binlog_Ignore_DB | Executed_Gtid_Set |
    +------------------+----------+--------------+------------------+-------------------+
    | mysql-bin.000005      120|              | mysql            |                   |
    +------------------+----------+--------------+------------------+-------------------+
    1 row in set (0.00 sec)
```
4.配置从库
```bash
    mysql> CHANGE MASTER TO 
    MASTER_HOST = '172.17.0.3',  
    MASTER_USER = 'repl_user', 
    MASTER_PASSWORD = 'repl_passwd',
    MASTER_PORT = 3307,
    MASTER_LOG_FILE='mysql-bin.000005',
    MASTER_LOG_POS=120,
    MASTER_RETRY_COUNT = 60,
    MASTER_HEARTBEAT_PERIOD = 10000; 
    
    # MASTER_LOG_FILE='mysql-bin.000005',#与主库File 保持一致
    # MASTER_LOG_POS=120 , #与主库Position 保持一致

```
5.启动从库
```bash
    slave start;
    Query OK, 0 rows affected (0.04 sec)

    如果想要同步所有库和表，在从mysql执行：
    
    STOP SLAVE SQL_THREAD; 
    CHANGE REPLICATION FILTER REPLICATE_DO_DB = (); 
    start SLAVE SQL_THREAD;
```
## 当主库为线上的时候， 采用 mysqldump 备份数据
```bash
    flush tables with read lock;
    show master status;
    #备份数据
    # 针对事务性引擎
    mysqldump -uroot -ptiger --all-database -e --single-transaction --flush-logs --max_allowed_packet=1048576 --net_buffer_length=16384 > /data/all_db.sql
    #恢复主库写锁
    unlock tables;


    从库
    mysql -uroot -p < /data/all_db.sql

```
