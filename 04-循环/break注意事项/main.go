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

	/*
		break的注意事项
		1）break语句出现在多层嵌套的语句块块时，可以通过标签指名要终止的是哪一层语句块
	*/

label2:
	for i := 0; i < 4; i++ {
		// label1: //设置一个标签
		for j := 0; j < 4; j++ {
			if j == 2 {
				// break label1
				break label2 //j=0 j=1
			}
			fmt.Println("j=", j)
		}
	}
}
