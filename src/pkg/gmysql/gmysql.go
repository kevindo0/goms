package gmysql

import (
	"strings"
	"time"
	"database/sql"
	"goms/pkg/setting"
	"goms/pkg/logging"
	_ "github.com/go-sql-driver/mysql"
)

var MysqlDb *sql.DB

func SetUp() error {
	path := strings.Join([]string{setting.MysqlSetting.Username, 
			":", setting.MysqlSetting.Password, 
			"@tcp(",setting.MysqlSetting.Ip, ":", setting.MysqlSetting.Port, ")/", 
			setting.MysqlSetting.Db, "?charset=utf8"}, "")
	MysqlDb, _ = sql.Open("mysql", path)
	// 设置数据库最大连接时间
    MysqlDb.SetConnMaxLifetime(setting.MysqlSetting.ConnMaxLifetime*time.Second)
    // 设置数据库最大连接数
    MysqlDb.SetMaxOpenConns(setting.MysqlSetting.MaxOpenConns)
    //设置数据库最大闲置连接数
    MysqlDb.SetMaxIdleConns(setting.MysqlSetting.MaxIdleConns)
    //验证连接
    if err := MysqlDb.Ping(); err != nil{
        logging.Error(err)
        return err
    }
    logging.Info("connnect success")
    return nil
}