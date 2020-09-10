package livetime

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

func LiveTimeEveryDay(db *gorm.DB, tlrIDs []int) {
	start, _ := time.Parse("2006-01-02", EveryDayStartTime)
	end, _ := time.Parse("2006-01-02", EveryDayENdTime)
	totalIntervalTime := 0
	// 总的观看人数
	number := 0
	// 总的观看时长
	totalSum := 0
	row := dataFileSave.EveryDaySheet.AddRow()
	for _, title := range []string{"日期", "观看时长", "观看人数", "每用户平均观看时长"} {
		nameCell := row.AddCell()
		nameCell.Value = title
	}
	for i := 0; i < 1000; i++ {
		s := start.AddDate(0, 0, i)
		e := start.AddDate(0, 0, i+1)
		if s.Unix() > end.Unix() {
			break
		}
		var interval struct {
			Sum int
		}
		var newdb *gorm.DB
		if len(tlrIDs) > 0 {
			newdb = db.Table("ziyuan_live_times").
				Where("created_at>? and created_at<? and live_room_id not in (?)", s, e, tlrIDs)
		} else {
			newdb = db.Table("ziyuan_live_times").
				Where("created_at>? and created_at<?", s, e)
		}
		newdb.Select("sum(`interval`) as sum").
			Find(&interval)

		if interval.Sum > 0 {
			totalIntervalTime += interval.Sum
			showDay := fmt.Sprintf("%d-%02d-%02d", s.Year(), s.Month(), s.Day())
			row := dataFileSave.EveryDaySheet.AddRow()
			row.AddCell().Value = showDay
			// 计算当日观看直播的总人数
			count := LiveRoomsWatchBetween(db, tlrIDs, s, e)
			average := interval.Sum / count
			number += count
			totalSum += interval.Sum
			// show := fmt.Sprintf("%s \t观看时长：%d(s) %s \t观看人数:%d \t每用户平均观看时长: %s", showDay, interval.Sum, Seconds2Time(interval.Sum), count, Seconds2TimeMinite(average))
			// row.AddCell().Value = fmt.Sprintf("%d(s) %s", interval.Sum, Seconds2Time(interval.Sum))
			row.AddCell().SetInt(interval.Sum)
			row.AddCell().SetInt(count)
			row.AddCell().Value = Seconds2TimeMinite(average)
		}
	}
	show := ""
	row = dataFileSave.EveryDaySheet.AddRow()
	if number > 0 {
		row = dataFileSave.EveryDaySheet.AddRow()
		res := totalSum / number
		row.AddCell().Value = ""
		row.AddCell().Value = fmt.Sprintf("总时长:%d", totalSum)
		row.AddCell().Value = fmt.Sprintf("人数:%d", number)
		row.AddCell().Value = fmt.Sprintf("平均后：%d(s), %s", res, Seconds2Time(res))
		show = fmt.Sprintf("根据每日数据汇总后，总的每用户平均观看时长：总时长%d(s), 人数:%d, 平均后：%d(s), %s", totalSum, number, res, Seconds2Time(res))
	} else {
		show = "没有人观看"
	}
	fmt.Println(show)
}

func LiveTimeRoomData(db *gorm.DB, tlrIDs []int) {
	type Data struct {
		LiveRoomID int
		Count      int
	}
	var data []Data
	err := db.Table("ziyuan_live_times").
		Where("live_room_id not in (?)", tlrIDs).
		Select("live_room_id, count(id) as count").
		Group("live_room_id").
		Order("count desc").
		Limit(30).
		Scan(&data).
		Error
	if err != nil {
		fmt.Sprintf("live room Data error", err)
		return
	}
	fmt.Println("观看次数最多的项目前30:")
	for _, d := range data {
		var name struct {
			Title string
		}
		db.Table("ziyuan_live_rooms").Where("id=?", d.LiveRoomID).Scan(&name)
		// user_id > 0
		count1 := 0
		db.Table("ziyuan_live_times").
			Where("live_room_id=? and user_id>0", d.LiveRoomID).
			Select("count(distinct(user_id))").
			Count(&count1)
		// user_id = 0
		count2 := 0
		db.Table("ziyuan_live_times").
			Where("live_room_id=? and user_id=0", d.LiveRoomID).
			Select("count(distinct(vid))").
			Count(&count2)
		count := count1 + count2

		var pv struct {
			Counter int
		}
		db.Table("ziyuan_live_data").
			Where("live_room_id=? and name='pv'", d.LiveRoomID).
			Scan(&pv)
		pvcount := pv.Counter
		if pv.Counter < d.Count+12 {
			pvcount = d.Count + 12
		}

		res := fmt.Sprintf("\t项目名称:%s, \t观看次数:%d, \tPV数:%d \t观看人数:%d", name.Title, d.Count, pvcount, count)
		fmt.Println(res)
	}
}
