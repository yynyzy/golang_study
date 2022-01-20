package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//生成一个随机数,我们需要给rand设置一个种子
	//time.Now().Unix() ：返回一个从 1970.01.01 的0分0秒的到现在的秒数
	rand.Seed(time.Now().Unix())

	n := rand.Intn(100) + 1 //[0 100]
	fmt.Println(n)

	//编写一个无限循环的控制，输出随机数，当生成99时，就会退出这个无限循环 break
	count := 0
	rand.Seed(time.Now().Unix())
	for {
		n := rand.Intn(100) + 1 //[0 100]
		count++
		if n == 99 {
			break
		}
	}
	fmt.Println("生成99一共用了", count)
}
