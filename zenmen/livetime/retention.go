package livetime

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

// 未区分是否是测试直播
func UserRetention(db *gorm.DB) {
	start, _ := time.Parse("2006-01-02", RetentionStartTime)
	end, _ := time.Parse("2006-01-02", RetentionEndTime)
	row := dataFileSave.RetentionSheet.AddRow()
	for _, title := range []string{"日期", "注册人数", "次留", "三留", "7留", "15留", "30留"} {
		nameCell := row.AddCell()
		nameCell.Value = title
	}
	for i := 0; i < 1000; i++ {
		s := start.AddDate(0, 0, i)
		e := start.AddDate(0, 0, i+1)
		if e.Unix() > end.Unix() {
			break
		}
		var userIDs []int
		err := db.Table("user").
			Where("registered between ? and ?", s, e).
			Pluck("user_id", &userIDs).
			Error
		if err != nil {
			panic(fmt.Errorf("user retention user ids %s", err))
		}
		showDay := fmt.Sprintf("%d-%02d-%02d", s.Year(), s.Month(), s.Day())
		// fmt.Println("showDay:", showDay, len(userIDs))
		row := dataFileSave.RetentionSheet.AddRow()
		row.AddCell().Value = showDay
		row.AddCell().SetInt(len(userIDs))
		if len(userIDs) > 0 {
			now := time.Now().Unix()
			for _, step := range []int{2, 3, 7, 15, 30} {
				end := s.AddDate(0, 0, step)
				if end.Unix() > now {
					break
				}
				r, c := Retention(db, e, end, userIDs)
				row.AddCell().Value = fmt.Sprintf("%s(%d)", r, c)
			}
		}
	}
}

// start 用户注册日的第二天，如查询用户的日期是2020-07-18，start是2020-07-19
// step 次留 3留、7留、15留、30留
func Retention(db *gorm.DB, start time.Time, end time.Time, userIDs []int) (string, int) {
	length := len(userIDs)
	if length == 0 {
		return "0.0", 0
	}
	// end := start.AddDate(0, 0, step-1)
	err := db.Table("ziyuan_user_actives").
		Where("user_id in (?) and created_time between ? and ?", userIDs, start, end).
		Select("user").
		Group("user_id").
		Error

	count := 0
	subQuery := db.Table("ziyuan_user_actives").
		Select("user_id").
		Where("created_at between ? and ?", start, end).
		Where("user_id in (?)", userIDs).
		Group("user_id").
		SubQuery()
	err = db.Raw("SELECT count(*) FROM ? as t", subQuery).
		Count(&count).Error
	if err != nil {
		panic(fmt.Errorf("retention query %s", err))
	}
	res := fmt.Sprintf("%0.4f", float64(count)/float64(length))
	return res, count
}
