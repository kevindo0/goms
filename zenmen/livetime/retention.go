package livetime

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

// 未区分是否是测试直播
func UserRetention(db *gorm.DB) {
	start := time.Date(2020, 7, 18, 0, 0, 0, 0, time.UTC)
	now := time.Now().Unix()
	for i := 0; i < 1000; i++ {
		s := start.AddDate(0, 0, i)
		e := start.AddDate(0, 0, i+1)
		if s.Unix() > now {
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
		// fmt.Println(s, e, userIDs, len(userIDs))
		showDay := fmt.Sprintf("注册日期: %d-%02d-%02d", s.Year(), s.Month(), s.Day())
		show := fmt.Sprintf("%s\t未发现新注册人员", showDay)
		if len(userIDs) > 0 {
			r_2 := Retention(db, e, 2, userIDs)
			r_3 := Retention(db, e, 3, userIDs)
			r_7 := Retention(db, e, 7, userIDs)
			r_15 := Retention(db, e, 15, userIDs)
			r_30 := Retention(db, e, 30, userIDs)
			show = fmt.Sprintf("%s \t注册人数: %d \t次留: %s \t三留: %s\t七留: %s \t15留: %s \t30留: %s", showDay, len(userIDs), r_2, r_3, r_7, r_15, r_30)
		}
		fmt.Println(show)
	}
}

// start 用户注册日的第二天，如查询用户的日期是2020-07-18，start是2020-07-19
// step 次留 3留、7留、15留、30留
// 次留时 step=2
func Retention(db *gorm.DB, start time.Time, step int, userIDs []int) string {
	length := len(userIDs)
	if length == 0 {
		return "0.0"
	}
	end := start.AddDate(0, 0, step-1)
	err := db.Table("ziyuan_live_times").
		Where("user_id in (?) and created_time between ? and end", userIDs, start, end).
		Select("user").
		Group("user_id").
		Error

	count := 0
	subQuery := db.Table("ziyuan_live_times").
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
	return res
}
