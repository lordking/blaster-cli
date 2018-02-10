Blaster-cli
=============

一个简单的blaster脚手架管理工具. blaster用于创建工程

## 1 编译环境准备

### 1.1 Go语言环境(Linux)

```bash
# 以root用户身份登录
# 下载
$ wget https://storage.googleapis.com/golang/go1.9.2.linux-amd64.tar.gz
$ tar xzvf go1.9.2.linux-amd64.tar.gz
$ mv go /usr/local

# 设置GO环境
$ vi /etc/profie.d/go_env.sh

# 添加如下代码
export GOROOT=/usr/local/go

PATH=$PATH:$GOROOT/bin
export PATH

# root用户身份操作，或切换成普通用户
$ su [用户]

# 设置GO PATH
$ vi ~/.bash_profile

# 添加如下代码
export GOPATH=$HOME/go-project

PATH=$PATH:$GOPATH/bin
export PATH

# 生效GO PATH配置
$ source ~/.bash_profile
```

### 1.2 依赖库

```shell
$ go get -u github.com/Sirupsen/logrus
$ go get -u github.com/spf13/viper
$ go get -u github.com/gin-gonic/gin
$ go get -u github.com/spf13/cobra
```



## 3 安装说明

```bash
$ git clone https://github.com/lordking/blaster-cli.git $GOPATH/src/github.com/lordking/blaster-cli
$ cd $GOPATH/src/github.com/lordking/blaster-cli/blaster-cli
$ godep go install

$ blaster-cli install
```

## 4 如何使用

### 打印当前版本号

Usage:
```bash
$ blaster-cli version
```

### 打印所有可使用的样例

Usage:
```bash
$ blaster-cli list
```

### 根据当前样例创建一个项目

创建的项目位于GOPATH/src下。

Usage:
```bash
$ blaster-cli new <template-name> <project-in-GOPATH>
```

Example:

```bash
$ blaster-cli new database-mongo mongo_test
```

