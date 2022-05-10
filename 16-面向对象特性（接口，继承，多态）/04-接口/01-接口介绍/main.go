package main

import "fmt"

type Usb interface {
	start()
	stop()
}
type phone struct {
}

func (p phone) start() {
	fmt.Println("手机Starting")
}
func (p phone) stop() {
	fmt.Println("手机Starting")
}

type camera struct {
}

func (c camera) start() {
	fmt.Println("照相机Starting")
}
func (c camera) stop() {
	fmt.Println("照相机Stoping")
}

type computer struct {
}

//接受一个Usb接口类型
//只要实现了Usb接口（实现了Usb接口声明的所有方法）
func (computer computer) Working(usb Usb) {
	usb.start()
	usb.stop()
}
func main() {
	computer := computer{}
	phone := phone{}
	camera := camera{}
	computer.Working(phone)
	computer.Working(camera)
	var _phone Usb = phone
	_phone.start()
	_phone.stop()
}
