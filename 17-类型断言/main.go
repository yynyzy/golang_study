package main

import "fmt"

type Point struct {
	x int
	y int
}

func main() {
	var a interface{}
	var point Point = Point{1, 2}
	a = point //oK
	//如何将a赋给一个Point变量?
	var b Point
	// b = a不可以
	b = a.(Point)
	fmt.Println(b)
}
