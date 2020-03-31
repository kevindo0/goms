package setting

import (
	"fmt"
	"time"
	"github.com/go-ini/ini"
)

type App struct {
	SavePath 	string
	LoggingFile string
}

var AppSetting = &App{}

type Mysql struct {
	Ip 			string
	Username 	string
	Password 	string
	Port		string
	Db			string
	ConnMaxLifetime time.Duration
	MaxOpenConns int
	MaxIdleConns int
}

var MysqlSetting = &Mysql{}

type Redis struct {
	Host string
	Password string 
	MaxIdle int
	MaxActive int
	IdleTimeout time.Duration
}

var RedisSetting = &Redis{}

type Mongo struct {
	Addrs []string
	Timeout time.Duration
	Database string
	PoolLimit int
}

var MongoSetting = &Mongo{}

var cfg *ini.File

func Setup() {
	var err error
	cfg, err = ini.Load("/etc/goms/config.ini")
	if err != nil {
		fmt.Println("cfg error:", err)
	}
	mapTo("app", AppSetting)
	mapTo("mysql", MysqlSetting)
	mapTo("redis", RedisSetting)
	mapTo("mongo", MongoSetting)
}

func mapTo(s string, v interface{}) {
	err := cfg.Section(s).MapTo(v)
	if err != nil {
		fmt.Println("error:", err)
	}
}