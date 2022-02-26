package main

import (
	"fmt"
	service "golang_study/CustomerManager/Service"
)

type customerView struct {
	key             string
	customerService *service.CustomerService
}

func (cv *customerView) MainMenu() {
table:
	for {
		fmt.Println("-------客户信息管理软件-------")
		fmt.Println("1      添加客户")
		fmt.Println("2      修改客户")
		fmt.Println("3      删除客户")
		fmt.Println("4      客户列表")
		fmt.Println("5      退出软件")
		fmt.Println("请选择(1-4):")
		fmt.Scanln(&cv.key)
		switch cv.key {
		case "1":
			cv.add()
		case "2":
			cv.add()
		case "3":
			cv.add()
		case "4":
			cv.List()
		case "5":
			res := cv.exit()
			if res {
				break table
			}
		default:
			fmt.Println("请输入正确的选项")
		}
	}
}

func (cv *customerView) add() {}

// func (cv *customerView) () {}
// func (cv *customerView) MainMenu() {}
func (cv *customerView) List() {
	customer := cv.customerService.List()
	fmt.Println("---------------客户列表--------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t手机\t邮箱")
	for i := 0; i < len(customer); i++ {
		fmt.Println(customer[i].GetInfo())
	}
	fmt.Println("-------------客户列表完成-------------")
}
func (cv *customerView) exit() bool {

	fmt.Println("你确认要退出软件吗? y/n")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		} else {
			fmt.Println("你的输入错误，请输入正确选项！")
		}
	}
	if choice == "y" {
		return true
	} else {
		return false
	}
}
func main() {
	customerView := customerView{
		key: "",
	}
	customerView.customerService = service.NewCustomerService()
	customerView.MainMenu()
}
