package livetime

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

// 新注册人数
func NewRegistered(db *gorm.DB) int {
	count := 0
	err := db.Table("user").
		Where("registered between ? and ?", StartTime, EndTime).
		Count(&count).Error
	if err != nil {
		panic(fmt.Errorf("new registered user %s", err))
	}
	row := dataFileSave.BasicSheet.AddRow()
	row.AddCell().Value = "新注册人数:"
	row.AddCell().SetInt(count)
	return count
}

// 测试直播间的ids
func TestLiveRoomID(db *gorm.DB) []int {
	var TestLiveIDs []int
	err := db.Table("ziyuan_live_rooms").
		Where("created_at>=? and title like ?", StartTime, "%测试%").
		Pluck("id", &TestLiveIDs).Error
	if err != nil {
		panic(fmt.Errorf("count %s", err))
	}
	return TestLiveIDs
}

// 累计上架直播数
func LiveRoomCountToC(db *gorm.DB, tlrIDs []int) int {
	count := 0
	var newdb *gorm.DB
	if len(tlrIDs) > 0 {
		newdb = db.Table("ziyuan_live_rooms").
			Where("`to`=0 and created_at between ? and ? and id not in (?)", StartTime, EndTime, tlrIDs)
	} else {
		newdb = db.Table("ziyuan_live_rooms").
			Where("`to`=0 and created_at between ? and ?", StartTime, EndTime)
	}
	err := newdb.Count(&count).Error
	if err != nil {
		panic(fmt.Errorf("live room count %s", err))
	}
	row := dataFileSave.BasicSheet.AddRow()
	row.AddCell().Value = "累计上架TOC直播数:"
	row.AddCell().SetInt(count)
	return count
}

// 累计上架直播数
func LiveRoomCountToB(db *gorm.DB, tlrIDs []int) int {
	count := 0
	var newdb *gorm.DB
	if len(tlrIDs) > 0 {
		newdb = db.Table("ziyuan_live_rooms").
			Where("`to`=1 and created_at between ? and ? and id not in (?)", StartTime, EndTime, tlrIDs)
	} else {
		newdb = db.Table("ziyuan_live_rooms").
			Where("`to`=1 and created_at between ? and ?", StartTime, EndTime)
	}
	err := newdb.Count(&count).Error
	if err != nil {
		panic(fmt.Errorf("live room count %s", err))
	}
	row := dataFileSave.BasicSheet.AddRow()
	row.AddCell().Value = "累计上架TOB直播数:"
	row.AddCell().SetInt(count)
	return count
}

// 计算累计观看人数
func Count(db *gorm.DB, tlrIDs []int) int {
	countLogin := 0
	countNot := 0
	countVisitorID := 0

	// user_id
	newdb := db.Table("ziyuan_live_times").
		Where("visitor_id is null and user_id>0")
	if len(tlrIDs) > 0 {
		newdb = newdb.Where("live_room_id not in (?)", tlrIDs)
	}
	err := newdb.Select("count(distinct(user_id))").
		Count(&countLogin).Error
	if err != nil {
		panic(fmt.Errorf("live room count login %s", err))
	}

	// vid
	newdb = db.Table("ziyuan_live_times").
		Where("visitor_id is null and user_id=0")
	if len(tlrIDs) > 0 {
		newdb = newdb.Where("live_room_id not in (?)", tlrIDs)
	}
	err = newdb.Select("count(distinct(vid))").
		Count(&countNot).Error
	if err != nil {
		panic(fmt.Errorf("live room count not login %s", err))
	}
	// visitor_id
	newdb = db.Table("ziyuan_live_times").
		Where("visitor_id is not null")
	if len(tlrIDs) > 0 {
		newdb = newdb.Where("live_room_id not in (?)", tlrIDs)
	}
	err = newdb.Select("count(distinct(visitor_id))").
		Count(&countVisitorID).Error
	if err != nil {
		panic(fmt.Errorf("live room count visitor_id %s", err))
	}

	count := countLogin + countNot + countVisitorID
	row := dataFileSave.BasicSheet.AddRow()
	row.AddCell().Value = "累计观看人数:"
	row.AddCell().SetInt(count)
	return count
}

// 上架新系统的直播累计播放时长
func LiveTimesTotalTime(db *gorm.DB, tlrIDs []int) int {
	type Result struct {
		Total int
	}
	results := Result{}

	newdb := db.Table("ziyuan_live_times").
		Select("sum(`interval`) as total")
	if len(tlrIDs) > 0 {
		newdb = newdb.Where("live_room_id not in (?)", tlrIDs)
	}
	err := newdb.Scan(&results).Error
	if err != nil {
		panic(fmt.Errorf("live times total time %s", err))
	}
	row := dataFileSave.BasicSheet.AddRow()
	row.AddCell().Value = "累计播放时长:"
	row.AddCell().Value = fmt.Sprintf("%d(s)    %s", results.Total, Seconds2Time(results.Total))
	return results.Total
}

// 所有用户累计观看次数
// 去除观看时长小于10s的
func LiveTimesTotalNumber(db *gorm.DB, tlrIDs []int) int {
	count := 0
	newdb := db.Table("ziyuan_live_times").
		Where("`interval` > 10")
	if len(tlrIDs) > 0 {
		newdb = newdb.Where("live_room_id not in (?)", tlrIDs)
	}
	err := newdb.Count(&count).
		Error
	if err != nil {
		panic(fmt.Errorf("live times total number %s", err))
	}
	row := dataFileSave.BasicSheet.AddRow()
	row.AddCell().Value = "所有用户累计观看次数:"
	row.AddCell().Value = fmt.Sprintf("%d 次", count)
	return count
}

// 所有用户累计观看项目总和
func LiveRoomsWatch(db *gorm.DB, tlrIDs []int) int {
	countLogin := 0
	countNot := 0
	countVisitorID := 0

	// user_id
	newdb := db.Table("ziyuan_live_times").
		Select("live_room_id, user_id").
		Where("visitor_id is null and user_id>0")
	if len(tlrIDs) > 0 {
		newdb = newdb.Where("live_room_id not in (?)", tlrIDs)
	}
	subQuery := newdb.Group("live_room_id,user_id").SubQuery()

	err := db.Raw("SELECT count(*) FROM ? as t", subQuery).
		Count(&countLogin).Error
	if err != nil {
		panic(fmt.Errorf("live room count login %s", err))
	}

	// vid
	newdb = db.Table("ziyuan_live_times").
		Select("live_room_id, vid").
		Where("visitor_id is null and user_id=0")
	if len(tlrIDs) > 0 {
		newdb = newdb.Where("live_room_id not in (?)", tlrIDs)
	}
	subQuery = newdb.Group("live_room_id,vid").SubQuery()
	err = db.Raw("SELECT count(*) FROM ? as t", subQuery).
		Count(&countNot).Error
	if err != nil {
		panic(fmt.Errorf("live room count not login %s", err))
	}
	// visitor_id
	newdb = db.Table("ziyuan_live_times").
		Select("live_room_id, visitor_id").
		Where("visitor_id is not null")
	if len(tlrIDs) > 0 {
		newdb = newdb.Where("live_room_id not in (?)", tlrIDs)
	}
	subQuery = newdb.Group("live_room_id,visitor_id").SubQuery()
	err = db.Raw("SELECT count(*) FROM ? as t", subQuery).
		Count(&countVisitorID).Error
	if err != nil {
		panic(fmt.Errorf("live room count visitor_id %s", err))
	}
	count := countLogin + countNot + countVisitorID
	row := dataFileSave.BasicSheet.AddRow()
	row.AddCell().Value = "所有用户累计观看项目总和数:"
	row.AddCell().SetInt(count)
	return count
}

// 所有用户累计观看项目总和
func LiveRoomsWatchBetween(db *gorm.DB, tlrIDs []int, start time.Time, end time.Time) int {
	countLogin := 0
	countNot := 0
	countVisitorID := 0
	// user_id
	newdb := db.Table("ziyuan_live_times").
		Select("live_room_id, user_id").
		Where("visitor_id is null and user_id>0").
		Where("created_at between ? and ?", start, end)
	if len(tlrIDs) > 0 {
		newdb = newdb.Where("live_room_id not in (?)", tlrIDs)
	}
	subQuery := newdb.Group("live_room_id,user_id").SubQuery()

	err := db.Raw("SELECT count(*) FROM ? as t", subQuery).
		Count(&countLogin).Error
	if err != nil {
		panic(fmt.Errorf("live room count login %s", err))
	}

	// vid
	newdb = db.Table("ziyuan_live_times").
		Select("live_room_id, vid").
		Where("visitor_id is null and user_id=0").
		Where("created_at between ? and ?", start, end)

	if len(tlrIDs) > 0 {
		newdb = newdb.Where("live_room_id not in (?)", tlrIDs)
	}
	subQuery = newdb.Group("live_room_id,vid").SubQuery()
	err = db.Raw("SELECT count(*) FROM ? as t", subQuery).
		Count(&countNot).Error
	if err != nil {
		panic(fmt.Errorf("live room count not login %s", err))
	}
	// visitor_id
	newdb = db.Table("ziyuan_live_times").
		Select("live_room_id, visitor_id").
		Where("visitor_id is not null").
		Where("created_at between ? and ?", start, end)

	if len(tlrIDs) > 0 {
		newdb = newdb.Where("live_room_id not in (?)", tlrIDs)
	}
	subQuery = newdb.Group("live_room_id,visitor_id").SubQuery()
	err = db.Raw("SELECT count(*) FROM ? as t", subQuery).
		Count(&countVisitorID).Error
	if err != nil {
		panic(fmt.Errorf("live room count visitor_id %s", err))
	}
	count := countLogin + countNot + countVisitorID
	return count
}
