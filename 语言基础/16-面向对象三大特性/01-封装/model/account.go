package model

import "fmt"

//定义一个结构体account
type account struct {
	accountNo string
	pwd       string
	balance   float64
}

//工厂模式的函数-构造函数
func NewAccount(accountNo string, pwd string, balance float64) *account {
	if len(accountNo) < 6 || len(accountNo) > 10 {
		fmt.Println("账号的长度不对...")
		return nil
	}
	if len(pwd) != 6 {
		fmt.Println("密码的长度不对...")
		return nil
	}
	if balance < 2 {

		fmt.Println("余额不足...")
		return nil
	}
	return &account{
		accountNo: accountNo,
		pwd:       pwd,
		balance:   balance,
	}
}
