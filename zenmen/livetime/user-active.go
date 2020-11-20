package livetime

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

// 活跃用户人数
func UserAciveEveryDay(db *gorm.DB) {
	start, _ := time.Parse("2006-01-02", UserActiveStartTime)
	end, _ := time.Parse("2006-01-02", UserActiveEndTime)
	row := dataFileSave.UserActiveSheet.AddRow()
	for _, title := range []string{"日期", "活跃人数"} {
		nameCell := row.AddCell()
		nameCell.Value = title
	}
	for i := 0; i < 1000; i++ {
		s := start.AddDate(0, 0, i)
		e := start.AddDate(0, 0, i+1)
		if e.Unix() > end.Unix() {
			break
		}
		date := s.Format("2006-01-02")
		count := 0
		countVisitorID := 0
		err := db.Table("ziyuan_user_actives").
			Where("visitor_id='' and date=?", date).
			Count(&count).Error
		if err != nil {
			fmt.Println("Error:", err)
		}
		err = db.Table("ziyuan_user_actives").
			Select("visitor_id").
			Where("visitor_id!='' and date=?", date).
			Group("visitor_id").
			Count(&countVisitorID).Error
		if err != nil {
			fmt.Println("Error:", err)
		}
		row := dataFileSave.UserActiveSheet.AddRow()
		row.AddCell().Value = date
		row.AddCell().SetInt(count + countVisitorID)
	}
}
