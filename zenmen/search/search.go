package search

import (
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

const (
	LiveItem   = "live"
	CourseItem = "course"
)

type LiveRoom struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Brief       string `json:"brief"`
	Description string `json:"description"`
	Tags        string `json:"tags"`
	UUID        string `json:"uuid"`
}

type Course struct {
	ID            uint   `json:"id"`
	Title         string `json:"title"`
	Brief         string `json:"brief"`
	Description   string `json:"description"`
	Tags          string `json:"tags"`
	SpecialColumn string `json:"special_column"`
	UUID          string `json:"uuid"`
}

func SearchInit() {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", User, Pwd, Host, Port, DB)
	db, err := gorm.Open("mysql", dns)
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	defer db.Close()
	fmt.Println("server is started...")
	LiveSearchInit(db)
	CourseInit(db)
}

func LiveSearchInit(db *gorm.DB) {
	rooms := []LiveRoom{}
	err := db.Table("ziyuan_live_rooms").
		Select("id,title,brief,description,tags,uuid").
		Where("public is true and status='online' and `to`=0").
		Scan(&rooms).Error
	if err != nil {
		panic(err)
	}
	for _, room := range rooms {
		names := []string{}
		if room.Tags != "" {
			names = append(names, room.Tags)
		}
		var lecturerIDs []int
		err = db.Table("ziyuan_lecturers").
			Where("live_room_id=? and user_id is not null", room.ID).
			Pluck("user_id", &lecturerIDs).Error
		if err != nil {
			fmt.Println("error:", room.ID, err)
			continue
		}
		if len(lecturerIDs) > 0 {
			realNames := []string{}
			err := db.Table("ziyuan_user_lecturers").
				Where("user_id in (?) and real_name is not null", lecturerIDs).
				Pluck("real_name", &realNames).
				Error
			if err != nil {
				fmt.Println("search user", err)
				continue
			}
			for _, realname := range realNames {
				if realname != "" {
					names = append(names, realname)
				}
			}
		}
		tags := ""
		if len(names) > 0 {
			tags = strings.Join(names, ",")
		}
		search := Search{
			Title:       room.Title,
			Brief:       room.Brief,
			Description: room.Description,
			Tags:        tags,
			UUID:        room.UUID,
		}
		var ids []int
		db.Table("ziyuan_searches").
			Where("item='live' and item_id=?", room.ID).
			Pluck("id", &ids)
		if len(ids) == 0 {
			search.Item = "live"
			search.ItemID = room.ID
			err := db.Table("ziyuan_searches").Create(&search).Error
			if err != nil {
				fmt.Println("error search create:", err)
			}
		} else {
			err := db.Table("ziyuan_searches").
				Where("item='live' and item_id=?", room.ID).
				Update(&search).Error
			if err != nil {
				fmt.Println("error search update:", err)
			}
		}
	}
}

func CourseInit(db *gorm.DB) {
	courses := []Course{}
	err := db.Table("ziyuan_courses").
		Select("id,title,brief,description,tags,special_column,uuid").
		Where("status='online'").
		Scan(&courses).Error
	if err != nil {
		panic(err)
	}
	for _, course := range courses {
		search := Search{
			Title:         course.Title,
			Brief:         course.Brief,
			Description:   course.Description,
			Tags:          course.Tags,
			SpecialColumn: course.SpecialColumn,
			UUID:          course.UUID,
		}
		var ids []int
		db.Table("ziyuan_searches").
			Where("item=? and item_id=?", CourseItem, course.ID).
			Pluck("id", &ids)
		if len(ids) == 0 {
			search.Item = CourseItem
			search.ItemID = course.ID
			err := db.Table("ziyuan_searches").Create(&search).Error
			if err != nil {
				fmt.Println("error search create:", err)
			}
		} else {
			err := db.Table("ziyuan_searches").
				Where("item=? and item_id=?", CourseItem, course.ID).
				Update(&search).Error
			if err != nil {
				fmt.Println("error search update:", err)
			}
		}
	}
}
