# file_transfer

---
[中文README](https://github.com/vearne/file_transfer/blob/master/README_zh.md)

### Overview
file_transfer is a very simple file transfer service. The principle is to upload the file to a temporary file server, and then you can get the file with address of the HTTP protocol. You or your partner can use wget or browser to download the file.


### Configuration file
file_transfer.yml

```
# bind address
listen_address: ":8080"
# the directory where the file is saved
upload_dir: "/tmp"
# URL prefix of the generated download address
url_prefix: "http://localhost:8080/download"

basic_auth:
  # enable basic_auth
  enabled: false
  username: "vearne"
  password: "happyft"
```


### Usage
#### Build
```
make build
```  
#### install 
```
make install
```  
default install directory
```
/usr/local/bin/
```

#### Start
```
file_transfer -c /tmp/file_transfer.yaml
```
If no configuration file is specified, 
the search order of the configuration file is

1)  Current work directory
2) /etc/
3) /etc/file_transfer

#### Stop
```
ps -ef| grep file_tran|head -n 1 |awk '{print $2}'|xargs kill
```

### Upload/Download Example
disenable basic auth
```
curl -F file=@tt.png http://localhost:8080/upload
```

```
wget http://localhost:8080/download/tt.png
```
enable basic auth
```
curl -F file=@tt.png http://localhost:8080/upload --user vearne:happyft
```
```
wget --http-user=vearne --http-password=happyft http://localhost:8080/download/tt.png
```
You can also download the file by visiting the following URL directly in your browser.
```
http://localhost:8080/download
```

### Warning
Basic auth is less secure and if it is used in a production environment，you may put  your files in  potential security hazard.
