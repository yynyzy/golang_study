package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("命令行参数有", len(os.Args))

	//遍历命令行参数
	for i, v := range os.Args {
		fmt.Printf("args[%v]=%v", i, v)
	}
}
