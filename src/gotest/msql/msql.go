package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func recordStats(db *sql.DB, name string, price float64) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return 
	}
	defer func() {
		switch err {
		case nil:
			tx.Commit()
		default:
			tx.Rollback()
		}
	}()
	if _, err = tx.Exec("update user set price = price + 1;"); err != nil {
		return 
	}
	sql1 := "insert into user (name, price) values (?, ?);"
	if _, err = tx.Exec(sql1, name, price); err != nil {
		return 
	}
	return
}

func main() {
	fmt.Println("a")
	db, err := sql.Open("mysql", "root:123456@/test")
	if err != nil {
		fmt.Println(err)
		return 
	}
	defer db.Close()
	if err = recordStats(db, "lihao", 77.9); err != nil {
		fmt.Println(err)
	}
}