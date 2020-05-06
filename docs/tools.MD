### 1.snmp tools
```
apt-get install snmp
```

### 2.[Cron](https://github.com/robfig/cron)

### 3. 工具
#### 1. goimports
```bash
# １首先确认sublime text安装了Gosublime
# 从 工具栏--- Preferences----可以看到 Browse Package(第一个就是），打开后，即找到了插件安装路径
# 将代码下载到package目录
git clone https://margo.sh/GoSublime

# 2. 下载goimports: 
go get code.google.com/p/go.tools/cmd/goimports
```
```json
// ３配置sublime text，配制方法如下：preferences->package settings->gosublime->setting-user中添加如下内容
{
    "fmt_enabled": true,
    "fmt_cmd": ["goimports"]
}
```
