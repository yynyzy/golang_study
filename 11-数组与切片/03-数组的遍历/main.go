package main

import (
	"fmt"
)

func way1() {
	var intArr [6]int = [6]int{2, 3, 4, 5, 6, 7}
	for i, v := range intArr {
		fmt.Printf("intArr[%d]的值为%v\n", i, v)
	}
}
func way2() {
	var intArr [6]int = [6]int{2, 3, 4, 5, 6, 7}
	for i := 0; i < len(intArr); i++ {
		fmt.Printf("intArr[%d]的值为%v\n", i, intArr[i])
	}
}
func main() {
	way1()
	fmt.Println("下面是way2")
	way2()
}
