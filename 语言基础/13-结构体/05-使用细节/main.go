package main

import (
	"fmt"
	"unsafe"
)

/*
	1.结构体的所有字段在内存中是连续的
*/

type Point struct {
	x int
	y int
}

type Rect struct {
	leftUp, rightDown Point
}

type Rect2 struct {
	leftUp, rightDown *Point
}

func main() {
	r1 := Rect{Point{1, 2}, Point{3, 4}}

	//r1有四个 int，在内存中连续分配
	fmt.Printf("r1.leftUp.x的地址=%p,r1.leftUp.y的地址=%p\nr1.rightDown.x的地址=%p,r1.rightDown.y的地址=%p\n", &r1.leftUp.x, &r1.leftUp.y, &r1.rightDown.x, &r1.rightDown.y)
	fmt.Printf("的字节数是%d\n", unsafe.Sizeof(r1.leftUp))
	r2 := Rect2{&Point{1, 2}, &Point{3, 4}}

	//r2有两个*Point类型，这个两个*Point类型的本身地址也是连续的;
	//但是他们指向的地址不一定是连续
	fmt.Printf("r2.leftUp的地址=%p\nr2.rightDown的地址=%p\n", &r2.leftUp, &r2.rightDown)
	//r2.leftUp的地址=0xc000046230    //r2.rightDown的地址=0xc000046238
}
