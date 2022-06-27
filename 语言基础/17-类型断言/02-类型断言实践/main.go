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
	fmt.Println(p.Name + "Starting")
}
func (p Phone) stop() {
	fmt.Println(p.Name + "Starting")
}
func (p Phone) call() {
	fmt.Println(p.Name + "正在打电话")
}

type Camera struct {
	Name string
}

func (c Camera) start() {
	fmt.Println(c.Name + "Starting")
}
func (c Camera) stop() {
	fmt.Println(c.Name + "Stoping")
}

type Computer struct {
}

func (computer Computer) Working(usb Usb) {
	usb.start()
	//如果usb是指向Phone结构体变量，则还需要调用call方法
	//类型断言..[注意体会]
	if phone, ok := usb.(Phone); ok == true {
		phone.call()
	}

	usb.stop()
}
func main() {
	//定义一个usb接口裁组,可以存放Phone和camera的结构体变量
	var usbArr [3]Usb
	usbArr[0] = Phone{"vivo"}
	usbArr[1] = Phone{"小米"}
	usbArr[2] = Camera{"尼康"}

	//Phone还有一个特有的方法call(),请遍历Usb数组，如果是Phone变量，
	//除了调用Usb接口声明的方法外，还需要调用Phone特有方法call =》类型断言

	var computer Computer
	for _, v := range usbArr {
		computer.Working(v)
	}
}
