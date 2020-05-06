### 一、安装
#### 1. GRPC
进入到第一个$GOTPATH目录下，新建google.golang.org目录，拉取golang在github上的镜像库：
```bash
cd /usr/local/go/path/src   
mkdir google.golang.org
cd google.golang.org/
git clone https://github.com/grpc/grpc-go grpc
```

而在grpc下有许多常用的包，例如：

* metadata：定义了grpc所支持的元数据结构，包中方法可以对MD进行获取和处理
* credentials：实现了grpc所支持的各种认证凭据，封装了客户端对服务端进行身份验证所需要的所有状态，并做出各种断言
* codes：定义了grpc使用的标准错误码，可通用

#### 2. Protoc Plugin
```bash
go get -u github.com/golang/protobuf/protoc-gen-go
```
[下载地址](https://github.com/protocolbuffers/protobuf/releases)
#### 3. Protocol Buffers v3

```bash
wget https://github.com/protocolbuffers/protobuf/releases/download/v3.10.0/protobuf-all-3.10.0.tar.gz
tar zxvf protobuf-all-3.10.0.tar.gz
cd protobuf-3.10.0/
./configure
make
make install
protoc --version
ldconfig
protoc --version
```
```bash
protoc --go_out=plugins=grpc:. *.proto
```
#### 4. Grpc-gateway
grpc-gateway是protoc的一个插件。它读取gRPC服务定义，并生成一个反向代理服务器，将RESTful JSON API转换为gRPC。此服务器是根据gRPC定义中的自定义选项生成的。
```
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
```

### 二、文章
1. [Grpc+Grpc Gateway实践一 介绍与环境安装](https://segmentfault.com/a/1190000013339403#articleHeader0)
2. [Golang gRPC实践 连载二 Hello gRPC](https://segmentfault.com/a/1190000007909829)

### 三、Example
```bash
cd goms/src/ggrpc
protoc --version    //libprotoc 3.10.0
protoc -I. --go_out=plugins=grpc:. proto/echo.proto
protoc -I. --grpc-gateway_out=logtostderr=true:. proto/echo.proto
ls proto 			//echo.pb.go  echo.pb.gw.go  echo.proto
go run server/main.go
go run client/main.go
go run gateway/main.go
curl -X POST 127.0.0.1:8080/v1/echo -d '{"value": "duc"}'
```

### 四、Error
1.  google/api/annotations.proto
```bash
ls /home/kk/go/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.11.3/third_party/googleapis
```