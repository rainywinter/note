# 使用 etcd 服务发现工具 的笔记


测试环境： Ubuntu  
etcd版本：3.2.13

### 下载安装

地址：https://github.com/coreos/etcd/releases/  
选择稳定版本下载解压后将etcd etcdctl 复制到 /usr/local/bin 下

### 运行

无参数启动： etcd

    默认使用2379端口为客户端提供通讯， 并使用端口2380来进行服务器间通讯

未客户端工具etcdctl 配置环境变量 ETCDCTL_API=3【 试试设置前后运行etcdctl version 对比区别， api版本不同 】

### 测试

    etcdctl put name yc  
    etcdctl get name