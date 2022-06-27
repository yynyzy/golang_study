package main

import (
	"fmt"
	"runtime"
)

func main() {
	//查看 cpu核数
	var cpuNum int = runtime.NumCPU()
	fmt.Printf("first cpuNum = %v\n", cpuNum)

	//设置使用多少个 cpu核 运行go程序
	runtime.GOMAXPROCS(cpuNum - 1)
	fmt.Printf("ok")
}
