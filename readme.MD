## 2. Go ms
### 2.1 go mod
```bash
go mod init goms
# 设置GOPROXY代理
go env -w GOPROXY=https://goproxy.cn,direct
# 设置GOPRIVATE来跳过私有库，比如常用的Gitlab或Gitee，中间使用逗号分隔
go env -w GOPRIVATE=*.gitlab.com,*.gitee.com
# replace 链接到本地
go mod edit -replace=gitlab.51y5.net/zaojiu/common@v0.0.0-20200413081649-969ce8d7dee8=/Users/duchao/zenmen/common
```

### 2.2 [gin hot reload](https://github.com/gravityblast/fresh)
```bash
# 1. install
go get github.com/pilu/fresh
# 2. 将 GOPATH/bin 目录添加到系统的 PATH 中。

# 3. 进入项目所在目录
cd /path/to/goms
# 4. 运行
fresh
# 5. 浏览器访问
http://localhost:8080
```

## 2. Study
1. [go gin example](https://github.com/EDDYCJY/go-gin-example.git)
2. [log recard logrus](https://github.com/sirupsen/logrus)
3. [log zap](https://github.com/uber-go/zap)
4. [golang runtime model](https://www.jianshu.com/p/84bac7932394)
5. [Next() 和 Abort() 的含义](https://blog.csdn.net/cyberspecter/article/details/100602552)

## 3. 依赖
```bash
go get gopkg.in/ini.v1  					#	https://github.com/go-ini/ini
go get -u github.com/gin-gonic/gin  		# 	https://gin-gonic.com/docs/
go get github.com/gomodule/redigo/redis		# 	https://github.com/gomodule/redigo
go get gopkg.in/mgo.v2						#	http://labix.org/mgo
go get github.com/robfig/cron/v3@v3.0.0		# 	github.com/robfig/cron
```

## 4. 工具安装
1. [GoSublime安装](https://www.jianshu.com/p/5a1e38a50aa3)
```bash
# 点击菜单: Preferences --> Browse Packages 打开Packages directory
git clone https://margo.sh/GoSublime
```

## 5. 版本管理（go mod）
版本号由三个部分组成: 主版本.副版本.修订号
```bash
# 可以创建一个库
go mod init goms
# 自动分析代码中的依赖，并获取相应版本的库
go build
# 表示安装特定版本的的依赖
go get package@version
# 移除go.mod中不再使用的依赖
go mod tidy
# 将依赖安装在仓库的vendor目录下
go mod verdor
go build  #命令会忽略掉vendor目录
go build -mod vendor   #命令会从当前目录的vendor目录下寻找依赖
```

## 6. Modules
1. [awesome go](https://awesome-go.com/)
2. [kafka sarama](https://github.com/Shopify/sarama)
3. [kafka sarama-cluster](https://github.com/bsm/sarama-cluster)
4. [web frame dotweb](https://github.com/devfeel/dotweb)
5. [链路追踪jaeger](https://github.com/jaegertracing/jaeger)

## 7. GDB
1. [delve工具](github.com/derekparker/delve/cmd/dlv)

## 8. Error
1. unrecognized import path "golang.org/x/crypto"
```bash
export GOPROXY=https://goproxy.io
```