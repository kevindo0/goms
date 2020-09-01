package livetime

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/tealeg/xlsx"
)

type DataSave struct {
	File           *xlsx.File
	RetentionSheet *xlsx.Sheet
}

var dataFileSave = &DataSave{}

func init() {
	fmt.Println("init")
	dataFileSave.File = xlsx.NewFile()
	dataFileSave.File.AddSheet("留存率")
}

func LiveTime() {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", User, Pwd, Host, Port, DB)
	db, err := gorm.Open("mysql", dns)
	if err != nil {
		fmt.Println(err)
	}
	// db.LogMode(true)
	defer db.Close()

	// // 新注册人数
	// newCounter := NewRegistered(db, StartTime)
	// fmt.Println("新注册人数：", newCounter)

	// // 获取测试直播的ids
	// testLiveRoomIDs := TestLiveRoomID(db)

	// // 累计上架直播数
	// liveRoomCountToc := LiveRoomCountToC(db, testLiveRoomIDs)
	// fmt.Println("累计上架TOC直播数: ", liveRoomCountToc)
	// // 累计上架直播数
	// liveRoomCountTob := LiveRoomCountToB(db, testLiveRoomIDs)
	// fmt.Println("累计上架TOB直播数: ", liveRoomCountTob)

	// // 计算累计观看人数
	// count := Count(db, testLiveRoomIDs)
	// fmt.Println("累计观看人数: ", count)

	// // 上架新系统的直播累计播放时长
	// liveTimeTotal := LiveTimesTotalTime(db, testLiveRoomIDs)
	// fmt.Println("累计播放时长: ", liveTimeTotal, Seconds2Time(liveTimeTotal))

	// // 所有用户累计观看次数
	// totalNumber := LiveTimesTotalNumber(db, testLiveRoomIDs)
	// fmt.Println("所有用户累计观看次数: ", totalNumber, " 次")

	// // 所有用户累计观看项目总和
	// liveRooms := LiveRoomsWatch(db, testLiveRoomIDs)
	// averageLiveRoomWatch := fmt.Sprintf("%0.2f", float64(liveRooms)/float64(count))
	// fmt.Println("所有用户累计观看项目总和数：", liveRooms)
	// fmt.Println("平均每用户在线时长:", Seconds2Time(liveTimeTotal/count))
	// fmt.Println("平均每用户观看项目数", averageLiveRoomWatch)

	// LiveTimeRoomData(db, testLiveRoomIDs)
	// LiveTimeUserData(db, testLiveRoomIDs)
	// LiveTimeEveryDay(db, testLiveRoomIDs)
	UserRetention(db)
}
