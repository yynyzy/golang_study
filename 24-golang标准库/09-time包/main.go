package main

import (
	"fmt"
	"time"
)

func test1() {
	now := time.Now() //获取当前时间
	// current time :2020-12-01 22:24:30.85736+8888 CST m=+0.080896031
	fmt.Printf("current time :%v\n", now)
	year := now.Year()     //年
	month := now.Month()   //月
	day := now.Day()       //日
	hour := now.Hour()     //小时
	minute := now.Minute() //分钟
	second := now.Second() //秒
	fmt.Printf("%d-%02d-%02d %82d :%82d :%82d\n", year, month, day, hour, minute, second)
	fmt.Printf("%T,%T,%T,%T,%T,%T,%T\n", now, year, month, day, hour, minute, second)
	// time.Time, int, time.Month, int, int, int, int
}
func main() {
	test1()
}
