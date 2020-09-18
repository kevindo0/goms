package livetime

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/tealeg/xlsx"
)

const (
	// OutFile = "output/livetime.xlsx"
	OutFile = "livetime.xlsx"
)

type DataSave struct {
	File            *xlsx.File
	BasicSheet      *xlsx.Sheet
	EveryDaySheet   *xlsx.Sheet
	RetentionSheet  *xlsx.Sheet
	LiveRoom        *xlsx.Sheet
	UserActiveSheet *xlsx.Sheet
}

func (d *DataSave) Save() {
	d.File.Save(OutFile)
}

var dataFileSave = &DataSave{}

func init() {
	fmt.Println("init")
	dataFileSave.File = xlsx.NewFile()
	dataFileSave.BasicSheet, _ = dataFileSave.File.AddSheet("基础数据")
	dataFileSave.EveryDaySheet, _ = dataFileSave.File.AddSheet("每日数据")
	dataFileSave.RetentionSheet, _ = dataFileSave.File.AddSheet("留存率")
	dataFileSave.LiveRoom, _ = dataFileSave.File.AddSheet("项目详情")
	dataFileSave.UserActiveSheet, _ = dataFileSave.File.AddSheet("用户活跃数")
}

func LiveTime() {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", User, Pwd, Host, Port, DB)
	db, err := gorm.Open("mysql", dns)
	if err != nil {
		fmt.Println(err)
	}
	// db.LogMode(true)
	defer db.Close()

	row := dataFileSave.BasicSheet.AddRow()
	row.AddCell().Value = "时间开始:"
	row.AddCell().Value = StartTime
	row.AddCell().Value = "时间结束(不包含):"
	row.AddCell().Value = EndTime

	// 新注册人数
	newCounter := NewRegistered(db)
	fmt.Println("新注册人数：", newCounter)

	// 获取测试直播的ids
	testLiveRoomIDs := TestLiveRoomID(db)

	// 累计上架直播数
	liveRoomCountToc := LiveRoomCountToC(db, testLiveRoomIDs)
	fmt.Println("累计上架TOC直播数: ", liveRoomCountToc)
	// 累计上架直播数
	liveRoomCountTob := LiveRoomCountToB(db, testLiveRoomIDs)
	fmt.Println("累计上架TOB直播数: ", liveRoomCountTob)

	// 计算累计观看人数
	count := Count(db, testLiveRoomIDs)
	fmt.Println("累计观看人数: ", count)

	// 上架新系统的直播累计播放时长
	liveTimeTotal := LiveTimesTotalTime(db, testLiveRoomIDs)
	fmt.Println("累计播放时长: ", liveTimeTotal, Seconds2Time(liveTimeTotal))

	// 所有用户累计观看次数
	totalNumber := LiveTimesTotalNumber(db, testLiveRoomIDs)
	fmt.Println("所有用户累计观看次数: ", totalNumber, " 次")

	// 所有用户累计观看项目总和
	liveRooms := LiveRoomsWatch(db, testLiveRoomIDs)
	averageLiveRoomWatch := fmt.Sprintf("%0.2f", float64(liveRooms)/float64(count))
	fmt.Println("所有用户累计观看项目总和数：", liveRooms)
	fmt.Println("平均每用户在线时长:", Seconds2Time(liveTimeTotal/count))
	fmt.Println("平均每用户观看项目数", averageLiveRoomWatch)
	row = dataFileSave.BasicSheet.AddRow()
	row.AddCell().Value = "平均每用户在线时长:"
	row.AddCell().Value = Seconds2Time(liveTimeTotal / count)

	row = dataFileSave.BasicSheet.AddRow()
	row.AddCell().Value = "平均每用户观看项目数:"
	row.AddCell().Value = averageLiveRoomWatch
	// LiveTimeRoomData(db, testLiveRoomIDs)
	// LiveTimeUserData(db, testLiveRoomIDs)
	LiveTimeEveryDay(db, testLiveRoomIDs)
	UserRetention(db)
	LiveRoomList(db, testLiveRoomIDs)
	UserAciveEveryDay(db)
	dataFileSave.Save()
}
