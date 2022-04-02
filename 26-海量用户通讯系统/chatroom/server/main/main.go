package main

import (
	"fmt"
	"golang_study/26-海量用户通讯系统/chatroom/server/model"
	"net"
	"time"
)

//处理和客户端的通信
func process(conn net.Conn) {
	//等待客户端发送信息
	defer conn.Close()
	//这里调用主控，创建一个process实例
	processor := &Processor{
		Conn: conn,
	}
	err := processor.processRecive()
	if err != nil {
		fmt.Println("客户端与服务端通讯协程错误 err=", err)
		return
	}

}

//此函数完成对 UserDao的初始化
func initUserDao() {
	//pool是一个全局变量，所以可以使用
	model.MyUserDao = model.NewUserDao(pool)
}
func main() {
	//在服务器一开始就创建一个 redis链接池
	initPool("localhost:6379", 16, 0, 300*time.Second)
	//注意要在初始化链接池后 再创建 UserDao实例
	initUserDao()
	fmt.Println("服务端[新的结构]监听8889端口")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	defer listen.Close()

	for {
		fmt.Println("服务端等待连接")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err=", err)
		}
		go process(conn)
	}
}
