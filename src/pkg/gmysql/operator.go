package gmysql

import (
	"fmt"
	"database/sql"
    _ "github.com/go-sql-driver/mysql"
)

type User struct {
    ID   int64          `db:"id"`
    Name sql.NullString `db:"name"`
    Age  int            `db:"age"`
}

//查询单行
func queryOne(DB *sql.DB) {
    user := new(User)
    row := DB.QueryRow("select * from users where id=?", 1)
    if err := row.Scan(&user.ID, &user.Name, &user.Age); err != nil {
        fmt.Printf("scan failed, err:%v", err)
        return
    }
    fmt.Println(*user)
}

//查询多行
func queryMulti(DB *sql.DB) {
    user := new(User)
    rows, err := DB.Query("select * from users where id > ?", 1)
    defer func() {
        if rows != nil {
            rows.Close()
        }
    }()
    if err != nil {
        fmt.Printf("Query failed,err:%v", err)
        return
    }
    for rows.Next() {
        err = rows.Scan(&user.ID, &user.Name, &user.Age)
        if err != nil {
            fmt.Printf("Scan failed,err:%v", err)
            return
        }
        fmt.Print(*user)
    }

}

//插入数据
func insertData(DB *sql.DB){
    result,err := DB.Exec("insert INTO users(name,age) values(?,?)","YDZ",23)
    if err != nil{
        fmt.Printf("Insert failed,err:%v",err)
        return
    }
    lastInsertID,err := result.LastInsertId()
    if err != nil {
        fmt.Printf("Get lastInsertID failed,err:%v",err)
        return
    }
    fmt.Println("LastInsertID:",lastInsertID)
    rowsaffected,err := result.RowsAffected()
    if err != nil {
        fmt.Printf("Get RowsAffected failed,err:%v",err)
        return
    }
    fmt.Println("RowsAffected:",rowsaffected)
}

//更新数据
func updateData(DB *sql.DB){
    result,err := DB.Exec("UPDATE users set age=? where id=?","30",3)
    if err != nil{
        fmt.Printf("Insert failed,err:%v",err)
        return
    }
    rowsaffected,err := result.RowsAffected()
    if err != nil {
        fmt.Printf("Get RowsAffected failed,err:%v",err)
        return
    }
    fmt.Println("RowsAffected:",rowsaffected)
}