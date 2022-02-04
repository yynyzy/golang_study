package main

import (
	"fmt"
)

func way1() {
	var intArr1 [6]int = [6]int{2, 3, 4, 5, 6, 7}
	for i, v := range intArr1 {
		fmt.Printf("intArr1[%d]的值为%v\n", i, v)
	}
}
func way2() {
	intArr2 := [...]string{"yyn", "yzy", "qq"}
	for i := 0; i < len(intArr2); i++ {
		fmt.Printf("intArr2[%d]的值为%v\n", i, intArr2[i])
	}
}
func main() {
	way1()
	fmt.Println("下面是way2")
	way2()
}
