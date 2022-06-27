package main

/*
基本介绍
变量(实例)具有多种形态。面向对象的第三大特征，在Go语言，多态特征是通过接口实现的。可以按照统一的接口来调用不同的实现。这时接口变量就呈现不同的形态。
*/
/*
接口体现多态的两种方式
1.多态参数
2.多态接口
*/
import "fmt"

type Usb interface {
	start()
	stop()
}
type Phone struct {
	Name string
}

func (p Phone) start() {
	fmt.Println("手机Starting")
}
func (p Phone) stop() {
	fmt.Println("手机Starting")
}

type Camera struct {
	Name string
}

func (c Camera) start() {
	fmt.Println("照相机Starting")
}
func (c Camera) stop() {
	fmt.Println("照相机Stoping")
}
func main() {
	//定义一个usb接口裁组,可以存放Phone和camera的结构体变量
	//这里就体现出多态数组
	var usbArr [3]Usb
	usbArr[0] = Phone{"vivo"}
	usbArr[1] = Phone{"小米"}
	usbArr[2] = Camera{"尼康"}
	fmt.Println(usbArr)
}
