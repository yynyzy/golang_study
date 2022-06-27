package main

import (
	"fmt"
	"runtime"
	"time"
)

func show() {

	for i := 0; i < 9; i++ {
		if i >= 5 {
			runtime.Goexit()
		}
		fmt.Printf("i=%v\n", i)
	}

}

func main() {

	go show()

	time.Sleep(time.Second * 2)
	fmt.Println("hello")

}
