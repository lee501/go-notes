### kafka搭建与使用

#### kafka简介

----
    - Producer 消息生产者，负责产生消息发送到kafka服务器上
    - Consumer 消息消费者，从kafka服务器读取消息
    - Consumer Group 消费者群组， 每个消费者可以划分为一个特定的群组
    - Topic 消息的身份标识，消费者可以根据topic读取特定的数据
    - Broker Kafka集群中包含的服务器
    

#### kafka配置

----
- ZooKeeper是一个开源的分布式协调框架
- Kafka使用ZooKeeper作为协调框架

### 启动服务

----
- 使用Kafka需先启动一个ZooKeeper服务
```text
    > bin/zookeeper-server-start.sh config/zookeeper.properties
```

        - 默认的ZooKeeper链接服务为2181端口
        ```
           # Zookeeper connection string (see zookeeper docs for details).
           # This is a comma separated host:port pairs, each corresponding to a zk
           # server. e.g. "127.0.0.1:3000,127.0.0.1:3001,127.0.0.1:3002".
           # You can also append an optional chroot string to the urls to specify the
           # root directory for all kafka znodes.
           zookeeper.connect=localhost:2181
        ```

- 然后启动kafka服务
```text
    > bin/kafka-server-start.sh config/server.properties
```
        - config/server.properties是Kafka的配置文件，可以用来监听host、port、broker。 
        - producer和consumer的监听端口为9092， 可以通过配置文件来修改
    

### Kafka如何使用

----
- 命令行
    - 创建Topic
        - 后续消息生产者和消息消费者才能针对性的发送和消费数据
        ```text
              >bin/kafka-topics.sh --create --zookeeper localhost:2181 --replication-factor 1 --topic first
        ```
    - 通过命令查看已经创建的Topic
        ```text
          > bin/kafka-topics.sh --list --zookeeper localhost:2181
          first的topic
        ```
    - 发送消息
        - 启动一个终端，执行producer脚本，会出现消息输入提示符
        ```text
          >bin/kafka-console-producer.sh --broker-list localhost:9092 --topic first
          > 输入信息
        ```
    - 消费消息
        - 另起一个终端，执行consumer脚本，消费者一直会处于监听状态，当从生产端输入消息，消费端就会接受到消息
        ```text
          > bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic first --from-beginning
          > 显示消息
        ```
- Python调用Kafka接口
    - 创建Producer脚本
        ```py
            import json
            from kafka import KafkaProducer
      
            #data传入的数据
            data = {"a": 1, "b": 2}
            producer = KafkaProducer(bootstrap_servers=["localhost:9092"], 
                                     api_version=(0, 10, 0), 
                                     value_serializer=lambda v: json.dumps(v).encode('utf-8'))
            for i in range(1000)
                producer.send("first", data, partition=0)
        ```
    - 创建消费者脚本
        ```py
            from kafka import KafkaConsumer
            ​
            ​
            consumer = KafkaConsumer('test', 
                                     group_id="group2", 
                                     bootstrap_servers=["localhost:9092"])
                                     
            for msg in consumer:
                print(msg.value)
            ​
            # 输出
            b'{"a": 1, "b": 2, "c": 3}'
            b'{"a": 1, "b": 2, "c": 3}'
            b'{"a": 1, "b": 2, "c": 3}'
            b'{"a": 1, "b": 2, "c": 3}'
            b'{"a": 1, "b": 2, "c": 3}'
            b'{"a": 1, "b": 2, "c": 3}'
        ```
