### 1. Install
1. [Console](https://www.consul.io/downloads.html)
```bash
unzip consul_1.6.1_linux_amd64.zip
# 启动Consul agent的开发模式
consul agent -dev
# 查看Consul集群的成员
consul memebers
```


### 2. 命令
```bash
# install protoc-gen-micro
go get github.com/micro/protoc-gen-micro
cd goms/gmicro
protoc --micro_out=. --go_out=. proto/greeter.proto
# micro是以go-micro框架为核心的微服务管理工具，通过它可以方便查看go-micro服务情况。
go get github.com/micro/micro
# List nodes
micro network nodes
# List services
micro network services
# List routes
micro network routes
# Peer graph
micro network graph
micro --help
```

### Error
1. cannot load github.com/hashicorp/consul/api: ambiguous import
```bash
# 修改go.mod，增加以下内容
replace github.com/hashicorp/consul => github.com/hashicorp/consul v1.5.2
```