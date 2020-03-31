package sstorage

import (
	"fmt"
	"time"
	"strings"
	"strconv"
	"goms/pkg/gmgo"
    "gopkg.in/mgo.v2/bson"
)

type User struct {
	Id bson.ObjectId `bson:"_id"`
	Username string `bson:"name"`
	Regtime int64 	`bson:"regtime"`
}

func (user User) ToString() string {
    return fmt.Sprintf("%#v", user)
}

func MongoStorage() string {
	// for i := 0; i < 100; i ++ {
	// 	go func(i int) {
	// 		add(i)
	// 	}(i)
	// }
	// time.Sleep(3*time.Second)
	find()
	return "h"
}

func add(i int) {
	user := new(User)
    user.Id = bson.NewObjectId()
    user.Username = strings.Join([]string{"lilei-", strconv.Itoa(i)}, "")
    user.Regtime = time.Now().Unix()
    err := gmgo.GMonDB.C("test").Insert(user)
    if err == nil {
        fmt.Println("插入成功")
    } else {
        fmt.Println(err.Error())
        defer panic(err)
    }
}

//查询
func find() {
    var users []User
    gmgo.GMonDB.C("test").Find(bson.M{"name": "lilei-5"}).All(&users)
    for _, value := range users {
        fmt.Println(value.ToString())
    }
    //根据ObjectId进行查询
    idStr := "5d8048f81d41c8c1dc04879a"
    objectId := bson.ObjectIdHex(idStr)
    user := new(User)
    gmgo.GMonDB.C("test").Find(bson.M{"_id": objectId}).One(user)
    fmt.Println(user)
}
