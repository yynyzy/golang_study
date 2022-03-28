package main

import "fmt"

var UserId int
var UserPwd string

func main() {
	loop := true
	var key int
	fmt.Println("------------欢迎来到聊天室系统-------------")
	for loop {
		fmt.Println("\t\t\t 1.进入聊天室")
		fmt.Println("\t\t\t 2.注册用户")
		fmt.Println("\t\t\t 3.退出系统")
		fmt.Println("\t\t\t请输入{1～3}进行选择")
		fmt.Scanln(&key)

		switch key {
		case 1:
			fmt.Println("登陆聊天室")
			loop = false
		case 2:
			fmt.Println("注册用户")
			loop = false
		case 3:
			fmt.Println("退出系统")
			loop = false
		}
	}
	if key == 1 {
		//说明用户要登陆，输入用户id与密码
		fmt.Println("请输入用户Id")
		fmt.Scanln(&UserId)
		fmt.Println("请输入用户密码")
		fmt.Scanln(&UserPwd)

		err := login(UserId, UserPwd)
		if err != nil {
			fmt.Println("用户登陆成功")
		} else {
			fmt.Println("用户登陆失败")
		}
	} else if key == 2 {
		fmt.Println("注册用户")
	} else {
		fmt.Println("退出系统")
	}
}
