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
	ID           uint   `json:"id"`
	Title        string `json:"title"`
	Brief        string `json:"brief"`
	Description  string `json:"description"`
	Tags         string `json:"tags"`
	Progress     int    `json:"progress"`
	ShowPlayback *bool  `json:"show_playback"`
}

type Course struct {
	ID            uint   `json:"id"`
	Title         string `json:"title"`
	Brief         string `json:"brief"`
	Description   string `json:"description"`
	Tags          string `json:"tags"`
	SpecialColumn string `json:"special_column"`
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
		Select("id,title,brief,description,tags,progress,show_playback").
		Where("public is true and status='online' and `to`=0 and title not like '%测试%'").
		Scan(&rooms).Error
	if err != nil {
		panic(err)
	}
	for _, room := range rooms {
		if room.Progress > 1 && (room.ShowPlayback == nil || !(*room.ShowPlayback)) {
			continue
		}
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

		var ids []int
		db.Table("ziyuan_searches").
			Where("item='live' and item_id=?", room.ID).
			Pluck("id", &ids)
		if len(ids) == 0 {
			search := Search{
				Item:        LiveItem,
				ItemID:      room.ID,
				Title:       room.Title,
				Brief:       room.Brief,
				Description: room.Description,
				Tags:        tags,
			}
			err := db.Table("ziyuan_searches").Create(&search).Error
			if err != nil {
				fmt.Println("error search create:", err)
			}
		} else {
			search := map[string]string{
				"title":       room.Title,
				"brief":       room.Brief,
				"description": room.Description,
				"tags":        tags,
			}
			err := db.Table("ziyuan_searches").
				Where("item='live' and item_id=?", room.ID).
				Updates(search).Error
			if err != nil {
				fmt.Println("error search update:", err)
			}
		}
	}
}

func CourseInit(db *gorm.DB) {
	courses := []Course{}
	err := db.Table("ziyuan_courses").
		Select("id,title,brief,description,tags,special_column").
		Where("status='online' and title not like '%测试%'").
		Scan(&courses).Error
	if err != nil {
		panic(err)
	}
	for _, course := range courses {
		var ids []int
		db.Table("ziyuan_searches").
			Where("item=? and item_id=?", CourseItem, course.ID).
			Pluck("id", &ids)
		if len(ids) == 0 {
			search := Search{
				Item:          CourseItem,
				ItemID:        course.ID,
				Title:         course.Title,
				Brief:         course.Brief,
				Description:   course.Description,
				Tags:          course.Tags,
				SpecialColumn: course.SpecialColumn,
			}
			err := db.Table("ziyuan_searches").Create(&search).Error
			if err != nil {
				fmt.Println("error search create:", err)
			}
		} else {
			search := map[string]string{
				"title":          course.Title,
				"brief":          course.Brief,
				"description":    course.Description,
				"tags":           course.Tags,
				"special_column": course.SpecialColumn,
			}
			err := db.Table("ziyuan_searches").
				Where("item=? and item_id=?", CourseItem, course.ID).
				Updates(search).Error
			if err != nil {
				fmt.Println("error search update:", err)
			}
		}
	}
}
