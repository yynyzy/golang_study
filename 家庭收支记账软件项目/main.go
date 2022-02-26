package main

import "fmt"

func main() {
	//声明一个变量，保存接收用户输入的选项
	key := ""
	//账户余额
	// balance := 10000.0
	//每次收支金额
	// money := 0.0
	//每次收支说明
	// note := ""
	//收支明细
	details := "收支\t账户金额\t收支金额\t说  明"
	//显示这个主菜单
table:
	for {
		fmt.Println("-------家庭收支记账软件-------")
		fmt.Println("      收支明细")
		fmt.Println("      登记收入")
		fmt.Println("      登记支出")
		fmt.Println("      退出软件")
		fmt.Println("请选择(1-4):")
		fmt.Scanln(&key)
		switch key {
		case "1":
			fmt.Println("进入收支明细")
			fmt.Println(details)

		case "2":
			fmt.Println("进入登记收入")

		case "3":
			fmt.Println("进入登记支出")

		case "4":
			fmt.Println("退出软件")
			break table
		default:
			fmt.Println("请输入正确的选项")
		}
	}
}
