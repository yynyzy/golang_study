package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)

	/*
		counter := 0
		for _ = range ticker.C {
			counter++
			if counter > 3 {
				ticker.Stop()
				break
			}
			fmt.Println("ticker...")
		}
	*/
	chanint := make(chan int)
	go func() {
		for _ = range ticker.C {
			select {
			case chanint <- 1:
			case chanint <- 2:
			case chanint <- 3:
			}
		}
	}()

	sum := 0
	for v := range chanint {
		fmt.Println("收到", v)
		sum += v
		if sum > 10 {
			break
		}
	}
}
