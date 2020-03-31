package gmgo

import (
	"fmt"
	"time"
	"gopkg.in/mgo.v2"
	"goms/pkg/setting"
	"goms/pkg/logging"
)

var GMonDB *mgo.Database
var GMonSess *mgo.Session

func SetUp() error {
	fmt.Println(len(setting.MongoSetting.Addrs), setting.MongoSetting)
	MongoDBDialInfo := &mgo.DialInfo{
		Addrs: setting.MongoSetting.Addrs,
		Timeout: setting.MongoSetting.Timeout * time.Second,
		Database: setting.MongoSetting.Database,
    	PoolLimit: setting.MongoSetting.PoolLimit,
	}
	GMonSess, err := mgo.DialWithInfo(MongoDBDialInfo)
	if err != nil {
		logging.Error(err)
		return err
	}
	logging.Info("mongo connect success")
	GMonDB = GMonSess.DB(setting.MongoSetting.Database)
	return nil
}