package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

//插入、更新和删除操作都使用Exec方法。
//func (db *DB) Exec(query string, args ...interface{}) (Result, error)

//更新
func updateData() {
	sql := "update user_tb1 set username=?, password=? where id=?"
	ret, err := db.Exec(sql, "kite2", "kite123", "2")
	if err != nil {
		fmt.Printf("更新失败, err:%v\n", err)
		return
	}
	rows, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("更新行失败, err:%v\n", err)
		return
	}
	fmt.Printf("更新成功, 更新的行数： %d.\n", rows)
}

var db *sql.DB

func initDB() (err error) {
	dsn := "root:yyn990902@tcp(127.0.0.1:3306)/go_db?charset=utf8mb4&parseTime=True"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	err2 := db.Ping()
	if err2 != nil {
		return err2
	}
	return nil
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Println("连接成功")
	}

	updateData()
}
