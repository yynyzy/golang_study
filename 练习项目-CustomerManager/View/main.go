package main

import (
	"fmt"
	model "golang_study/练习项目-CustomerManager/Model"
	service "golang_study/练习项目-CustomerManager/Service"
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
			cv.delete()
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

// func (cv *customerView) () {}

func (cv *customerView) List() {
	customer := cv.customerService.List()
	fmt.Println("---------------客户列表--------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t手机\t邮箱")
	for i := 0; i < len(customer); i++ {
		fmt.Println(customer[i].GetInfo())
	}
	fmt.Println("-------------客户列表完成-------------")
}
func (cv *customerView) add() {
	fmt.Println("---------------添加客户---------------")
	fmt.Println("请输入名字")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("请输入性别")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("请输入年龄")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("请输入电话")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("请输入邮箱 ")
	email := ""
	fmt.Scanln(&email)
	customer := model.NewCustomer2(name, gender, age, phone, email)
	if cv.customerService.Add(customer) {
		fmt.Println("---------------添加成功---------------")
	} else {
		fmt.Println("---------------添加失败---------------")

	}
}
func (cv *customerView) delete() {
	fmt.Println("-----------------删除客户------------------")
	fmt.Println("请输入删除客户的编号")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return //放弃
	}
	fmt.Println("你确认要删除吗? y/n")
	choice := ""
	fmt.Scanln(&choice)
	if choice == "y" || choice == "Y" {
		if cv.customerService.Delete(id) {
			fmt.Println("删除成功")
		} else {
			fmt.Println("删除失败")
		}
	}
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
