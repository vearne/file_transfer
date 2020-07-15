# file_transfer

---
[English README](https://github.com/vearne/file_transfer/blob/master/README.md)

### 概述
file_transfer是一个非常简单的文件中转服务
其原理是将文件上传到临时的文件服务器，然后你可以得到HTTP协议的文件下载地址，你或你的伙伴可以使用wget或者浏览器来下载这个文件。


### 配置文件
file_transfer.yml

```
# 服务地址
listen_address: ":8080"
# 文件上传地址
upload_dir: "/tmp"
# 用于生成下载地址
url_prefix: "http://localhost:8080/download"

basic_auth:
  # 是否启用basic_auth
  # enable basic_auth
  enabled: false
  username: "vearne"
  password: "happyft"
```

### 使用
#### 编译
```
make build
```  
#### 安装 
```
make install
```  
默认安装到
```
/usr/local/bin/
```

#### 启动
```
file_transfer -c /tmp/file_transfer.yaml
```
如果不指定配置文件，配置文件的搜索顺序为

* 当前目录
* /etc/
* /etc/file_transfer

#### 停止
直接kill掉即可

### 上传/下载(样例)
不使用basic auth
```
curl -F file=@tt.png http://localhost:8080/upload 
```

```
wget http://localhost:8080/download/tt.png
```
使用basic auth
```
curl -F file=@tt.png http://localhost:8080/upload --user vearne:happyft
```
```
wget --http-user=vearne --http-password=happyft http://localhost:8080/download/tt.png 
```
你也可以直接在浏览器中访问以下地址，来下载文件。
```
http://localhost:8080/download
```

### 注意
basic auth 安全强度较差，如果用在生产环境有一定的风险，请慎重使用。
