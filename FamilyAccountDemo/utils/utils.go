package utils

import "fmt"

type FamilyAcccount struct {
	//声明一个变量，保存接收用户输入的选项
	key string
	//账户余额
	balance float64
	//每次收支金额
	money float64
	//每次收支说明
	note string
	//收支明细
	details string
	// := "收支\t账户金额\t收支金额\t说  明"

}

func NewFamilyAcccount() *FamilyAcccount {
	return &FamilyAcccount{
		key:     "",
		balance: 1000.0,
		money:   0,
		note:    "",
		details: "收支\t账户金额\t收支金额\t说  明",
	}
}
func (this *FamilyAcccount) MainMenu() {
table:
	for {
		fmt.Println("-------家庭收支记账软件-------")
		fmt.Println("      收支明细")
		fmt.Println("      登记收入")
		fmt.Println("      登记支出")
		fmt.Println("      退出软件")
		fmt.Println("请选择(1-4):")
		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.showDetail()
		case "2":
			this.income()
		case "3":
			this.pay()
		case "4":
			res := this.exit()
			if res {
				break table
			}
		default:
			fmt.Println("请输入正确的选项")
		}
	}
}

func (this *FamilyAcccount) showDetail() {
	fmt.Println(this.details)
}
func (this *FamilyAcccount) income() {
	fmt.Println("请输入本次收入金额")
	fmt.Scanln(&this.money)
	this.balance += this.money
	fmt.Println("请输入收支说明")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n收入\t 	%v\t		%v\t      %v", this.balance, this.money, this.note)

}

func (this *FamilyAcccount) pay() {
	fmt.Println("请输入本次支出金额")
	fmt.Scanln(&this.money)
	if this.money > this.balance {
		fmt.Println("账户余额不足")
		return
	}
	this.balance -= this.money
	fmt.Println("请输入收支说明")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n收入\t %v\t %v\t %v", this.balance, this.money, this.note)
}

func (this *FamilyAcccount) exit() bool {
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
