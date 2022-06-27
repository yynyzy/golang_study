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

//获取时间戳
func test2() {
	now := time.Now()
	//当前时间戳 TimeStamp type :int64，TimeStamp : 1654507225
	fmt.Printf("TimeStamp type:%T,TimeStamp: %v", now.Unix(), now.Unix())
}

func test3() {
	//使用time.Tick(时间间隔)来设置定时器，定时器的本质上是一个通道(channel)。
	ticker := time.Tick(time.Second)
	// 等价于 ticker2 := time.NewTicker(time.Second)
	for i := range ticker {
		fmt.Println(i)
	}
}

func test4() {
	t := time.Now()
	fmt.Println(t.Format("2006/01/0215:04"))

}
func main() {
	// test1()
	// test2()
	// test3()
	test4()
}
