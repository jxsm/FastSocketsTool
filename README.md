# FastSocketsTool

---

<div align="center">
    <img src="https://img.shields.io/github/go-mod/go-version/jxsm/FastSocketsTool" alt="go" />
    <img src="https://img.shields.io/github/license/jxsm/FastSocketsTool" alt="license">
    <img src="https://img.shields.io/github/last-commit/jxsm/FastSocketsTool" alt="commit">
</div>

<p align="center">
  English|
<a href="doc/markdown/chinese.md">简体中文</a>
</p>

<p align="center">
    <a href="doc/markdown/Acknowledgments.md">Acknowledgments</a>
</p>

---
## Introducing
FastSocketsTool is a Socket utility based on the Go language, offering a range of 
common Socket functionalities, including the transmission and reception of 
data, with the added capability to specify encoding methods for both sending and receiving operations.


## Compile

```shell
git clone https://github.com/jxsm/FastSocketsTool.git

// Enter the project directory.
cd FastSocketsTool

// Install dependencies.
go mod tidy

// compile
go build cmd/fst.go
```

## Use
```shell
./fst -h 127.0.0.1 -p 9000
```
<img src="doc/img/SimpleExample.png" alt="SimpleExample">

## Parameters
>Server mode is under development and currently only client mode is supported.

| Parameter | Instructions                                 | Default value |
|-----------|----------------------------------------------|---------------|
| -h        | Specify the IP address of the server         |               |
| -p        | Specify the port number of the server        |               |
| -e        | Specify the encoding method for transmission | UTF-8         |
| -re       | Specify the encoding method for reception.   | UTF-8         |
| -u        | UDP mode                                     | false         |
| -s        | Server mode                                  | false         |