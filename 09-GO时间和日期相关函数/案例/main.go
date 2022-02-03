package main

import (
	"fmt"
	"strconv"
	"time"
)

func test3() {
	str := ""
	for i := 0; i < 100000; i++ {
		str += "hello" + strconv.Itoa(i)
	}
}
func main() {
	//在执行test3前，先获取到当前的unix时间戳
	start := time.Now().Unix()
	test3()
	end := time.Now().Unix()
	fmt.Println("执行testo3()耗费时间为‰秒\n", end-start)
}
