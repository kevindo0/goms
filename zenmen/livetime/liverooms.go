package livetime

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type LiveRoomOrm struct {
	ID         int
	Title      string
	Created_at string
}

type LiveRoomPv struct {
	Counter int
}

func LiveRoomList(db *gorm.DB, tlrIDs []int) {
	row := dataFileSave.LiveRoom.AddRow()
	for _, title := range []string{"项目", "创建时间", "pv", "直播观看人数"} {
		nameCell := row.AddCell()
		nameCell.Value = title
	}
	var newdb *gorm.DB
	if len(tlrIDs) > 0 {
		newdb = db.Table("ziyuan_live_rooms").
			Where("created_at between ? and ? and id not in (?)", StartTime, EndTime, tlrIDs)
	} else {
		newdb = db.Table("ziyuan_live_rooms").
			Where("created_at between ? and ?", StartTime, EndTime)
	}
	var rooms []LiveRoomOrm
	err := newdb.Where("deleted_at is null").
		Select("id,title,created_at").Scan(&rooms).Error
	if err != nil {
		fmt.Println("live rooms", err)
	}
	for _, room := range rooms {
		row := dataFileSave.LiveRoom.AddRow()
		pv := GetLiveRoomPV(db, room.ID)
		count := GetLiveRoomCount(db, room.ID)
		row.AddCell().Value = room.Title
		row.AddCell().Value = room.Created_at
		row.AddCell().SetInt(pv)
		row.AddCell().SetInt(count)
	}
}

func GetLiveRoomPV(db *gorm.DB, liveRoomID int) int {
	var pv LiveRoomPv
	err := db.Table("ziyuan_live_data").
		Select("counter").
		Where("live_room_id=? and name='pv'", liveRoomID).
		Scan(&pv).Error
	if err != nil {
		fmt.Println("pv query", err)
	}
	return pv.Counter
}

// 计算累计观看人数
func GetLiveRoomCount(db *gorm.DB, liveRoomID int) int {
	countLogin := 0
	countNot := 0

	// 按user_id进行区分
	err := db.Table("ziyuan_live_times").
		Where("user_id>0 and live_room_id=?", liveRoomID).
		Select("count(distinct(user_id))").
		Count(&countLogin).Error
	if err != nil {
		panic(fmt.Errorf("live room count login %s", err))
	}
	// 按vid进行区分
	err = db.Table("ziyuan_live_times").
		Where("user_id=0 and live_room_id=?", liveRoomID).
		Select("count(distinct(vid))").
		Count(&countNot).Error
	if err != nil {
		panic(fmt.Errorf("live room count login %s", err))
	}
	count := countLogin + countNot
	return count
}
