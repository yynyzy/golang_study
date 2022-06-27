package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

/*
插入、更新和删除操作都使用Exec方法。
func (db *DB) Exec(query string, args ...interface{}) (Result, error)
*/

func insertData() {
	sqlStr := "insert into user_tb1(username,password) values (?,?)"
	r, err := db.Exec(sqlStr, "张三", "zs123")
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	theID, err := r.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
}

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

	insertData()
}
