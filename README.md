## 使用

`docker-compose up` 启动两个节点

集群信息 为： http://node1:8081,http://node2:8082,

将第三个节点信息交给集群：

```bash
 curl -L http://127.0.0.1:9901/3 -XPOST -d http://node3:8083
```

将 docker-compose 中 node3 的配置信息更新到集群：
在新的bash中执行docker-compose up，可以保证node3的启动和集群是在同一网络环境下


test 文件中含有压入数据脚本

```
 curl -H "Content-Type: application/json" -X POST -d '{  "Key": "'"${rand_key}"'", "Value": "'"${rand_val}"'"  }' "localhost:8072/kv"
```

注意： 后加入的集群拥有全部的 集群信息

RaftCluster = "http://node1:8081,http://node2:8082,http://node3:8083"

最开始启动几个节点，则RaftCluster中 地址有几个，按照逗号分隔。

以node1为例：

addr： 为对外服务用于对数据的增删改查
RaftPort： 为raft集群内部通信用
RaftCluster: 为集群内部通信用






## 分布式编程题
实现一个高可用，多副本，强一致，持久化的KV存储系统。实现语言为`Golang`。

本项目提供了整体的框架，答题者按照这个框架完善`pkg/store`，实现自己的store。

答题者可以使用`make docker`来编译打包镜像，使用`docker-compose up`来启动3副本测试。

### 1. 要求
* 对外客户端接口使用HTTP接口，实现`GET`, `SET`, `DELETE`语义
* 强一致协议使用raft，借助etcd的raft来构建
* 持久化使用[Pebble](https://github.com/cockroachdb/pebble)来构建
* 数据采用3副本存储
* Raft-Log需要删除，不能无限制存储，删除的时机和逻辑需要答题者自己决定

### 2. 如何测试
收到答题者的作品时，我们会按照以下的流程进行测试

* 启动2个节点
* 使用一个客户端以 1000/s 的速率持续的写入KV数据5分钟，并且记录服务端返回成功写入的数据
* 数据写入2分钟后，启动第三个节点
* 滚动重启3个节点
* 停止写入客户端
* 检查数据是否正确

### 3. 如何提交

调试完毕后，发送代码到邮件 zx AT matrixorigin DOT cn 

**WARNING: DO NOT DIRECTLY FORK THIS REPO. DO NOT PUSH PROJECT SOLUTIONS PUBLICLY.**
