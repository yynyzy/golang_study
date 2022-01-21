package main

import "fmt"

func main() {
	//goto 可以无条件的执行到程序中指定的行

	/*
		goto label
		...
		label statement
	*/
	var n int = 10
	fmt.Println("1")
	if n > 11 {
		goto label
	}
	fmt.Println("2")
	fmt.Println("3")
label:
	fmt.Println("4")
	fmt.Println("5")
}
