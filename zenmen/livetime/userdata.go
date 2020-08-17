package livetime

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// 根据用户观看时间记录进行排序
// 取前30名
func LiveTimeUserData(db *gorm.DB, tlrIDs []int) {
	type Data struct {
		UserID int
		Sum    int
	}
	var data []Data
	err := db.Table("ziyuan_live_times").
		Where("user_id>0 and live_room_id not in (?)", tlrIDs).
		Select("user_id, sum(`interval`) as sum").
		Group("user_id").
		Order("sum desc").
		Limit(30).
		Scan(&data).
		Error
	if err != nil {
		panic(fmt.Errorf("live time user data %s", err))
	}
	fmt.Println("观看时长最长的用户前30:")
	for _, d := range data {
		var user struct {
			NickName string
		}
		fmt.Print("user_id:", d.UserID)
		db.Table("user").Where("user_id=?", d.UserID).Scan(&user)
		count := 0
		db.Table("ziyuan_live_times").
			Where("user_id=?", d.UserID).
			Select("count(distinct(live_room_id))").
			Count(&count)
		res := fmt.Sprintf("\t昵称:%s, \t累计观看时长:%s, \t观看项目数:%d", user.NickName, Seconds2Time(d.Sum), count)
		fmt.Println(res)
	}
}
