package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

//删除数据
//插入、更新和删除操作都使用Exec方法。
//func (db *DB) Exec(query string, args ...interface{}) (Result, error)
func delData() {
	sql := "delete from user_tb1 where id =?"
	ret, err := db.Exec(sql, "1")
	if err != nil {
		fmt.Printf("删除失败, err:%v\n", err)
		return
	}
	rows, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("删除行失败, err:%v\n", err)
		return
	}
	fmt.Printf("删除成功, 删除的行数： %d.\n", rows)
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

	delData()
}
