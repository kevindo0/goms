package example_test

import (
	"fmt"
	"math/rand"

	"github.com/8treenet/gcache"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db          *gorm.DB
	cachePlugin gcache.Plugin
)

type TestUser struct {
	gorm.Model
	UserName string `gorm:"size:32"`
	Password string `gorm:"size:32"`
	Age      int
	Status   int
}

type TestEmail struct {
	gorm.Model
	TypeID     int
	Subscribed bool
	TestUserID int
}

func init() {
	var e error
	addr := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local"
	db, e = gorm.Open("mysql", addr)
	if e != nil {
		panic(e)
	}
	db.AutoMigrate(&TestUser{})
	db.AutoMigrate(&TestEmail{})

	opt := gcache.DefaultOption{}
	opt.Expires = 300              //缓存时间，默认60秒。范围 30-900
	opt.Level = gcache.LevelSearch //缓存级别，默认LevelSearch。LevelDisable:关闭缓存，LevelModel:模型缓存， LevelSearch:查询缓存
	opt.AsyncWrite = false         //异步缓存更新, 默认false。 insert update delete 成功后是否异步更新缓存
	opt.PenetrationSafe = false    //开启防穿透, 默认false。

	//缓存中间件 注入到Gorm
	cachePlugin = gcache.AttachDB(db, &opt, &gcache.RedisOption{Addr: "localhost:6379"})

	InitData()
	//开启Debug，查看日志
	db.LogMode(true)
	cachePlugin.Debug()
}

func InitData() {
	cachePlugin.FlushDB()
	db.Exec("truncate test_users")
	db.Exec("truncate test_emails")
	for index := 1; index < 21; index++ {
		user := &TestUser{}
		user.UserName = fmt.Sprintf("%s_%d", "name", index)
		user.Password = fmt.Sprintf("%s_%d", "password", index)
		user.Age = 20 + index
		user.Status = rand.Intn(3)
		db.Save(user)

		email := &TestEmail{}
		email.TypeID = index
		email.TestUserID = index
		db.Save(email)
	}
}
