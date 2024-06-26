# FastSocketsTool

---
<div align="center">
    <img src="https://img.shields.io/github/go-mod/go-version/jxsm/FastSocketsTool" alt="go" />
    <img src="https://img.shields.io/github/license/jxsm/FastSocketsTool" alt="license">
    <img src="https://img.shields.io/github/last-commit/jxsm/FastSocketsTool" alt="commit">
</div>

<p align="center">
  简体中文|
<a href="../../README.md">English</a>
</p>

<p align="center">
    <a href="Acknowledgments.md">致谢</a>
</p>

---
## 介绍

FastSocketsTool 是一个基于Go语言的Socket工具，它提供了一些常用的Socket功能，
例如发送和接收数据,并且可以指定其发送和接收时使用的编码方式

## 编译

```shell
git clone https://github.com/jxsm/FastSocketsTool.git

# 进入项目目录
cd FastSocketsTool

# 安装依赖
go mod tidy

# 编译
go build cmd/fst.go
```

## 使用
<img src="../img/client.png" alt="client">

## 服务器模式
```shell
./fst -h 127.0.0.1 -p 9000 -s
```
<img src="../img/server.png" alt="server">

当你在使用服务器模式时可以输入`list`来查看当前所有连接的客户端,使用`exit`来退出服务器模式

在默认情况下你在服务器模式下发送的数据是群发的,及每个连接者都可以收到你发送的数据,
但你可以在输入时使用`ip@内容`来指定发送给指定客户端的数据,例如:

```
127.0.0.1@你好
```

如果你需要发送如`list`,`exit`,`@`这种标识字符给所有人,则可以在开头输入`@`如:

```shell
# 输入
@@大家好
@list
# 所有连接者都会收到
@大家好
list
```


## 参数

| 参数  | 说明         | 默认值   |
|-----|------------|-------|
| -h  | 指定服务器的IP地址 |       |
| -p  | 指定服务器的端口号  |       |
| -e  | 指定发送编码方式   | UTF-8 |
| -re | 指定接收编码方式   | UTF-8 |
| -u  | udp模式      | false |
| -s  | 服务器模式      | false |
| -6  | ipv6模式     | false |
