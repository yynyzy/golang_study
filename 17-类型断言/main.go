package main

import "fmt"

/*
	类型断言，由于接口是一般类型，不知道具体类型，如果要转成具体类型，就需要用类型断言，具体的如下:

*/
type Point struct {
	x int
	y int
}

func main() {
	var a interface{} //空接口，可以接收任意类型
	var point Point = Point{1, 2}
	a = point //oK
	//如何将a赋给一个Point变量?
	var b Point
	// b = a不可以
	b = a.(Point)
	fmt.Println(b)

	//*2.类型断言(带检测的)
	var x interface{}
	var b2 float32 = 1.1
	x = b2

	//类型断言(带检测的)
	y, ok := x.(float64)
	if ok == true {
		fmt.Println("convert success")
		fmt.Printf("y的类型是%T值是=%v", y, y)
	} else {
		fmt.Println(" convert fail")
	}
	fmt.Println("继续执行...")

}
