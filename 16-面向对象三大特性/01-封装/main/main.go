package main

/*
封装(encapsulation)就是把抽象出的字段和对字段的操作封装在一起,数据被保护在内部,程序的其它包只有通过被授权的操作(方法),才能对字段进行操作。

封装的实现步骤
1)将结构体、字段(属性)的首字母小写(不能导出了，其它包不能使用，类似 private
2)给结构体所在包提供一个工厂模式的函数，首字母大写。类似一个构造函数
3)提供一个首字母大写的set方法(类似其它语言的public)，用于对属性判断并赋值
func (var结构体类型名)SetXxx(参数列表) (返回值列表){
	//加入数据验证的业务逻辑
	var.字段=参数
}
提供一个首字母大写的Get方法(类似其它语言的public)，用于获取属性的值
func (var结构体类型名) GetXxx(){
		return var.age
}

!特别说明。在Golang开发中并没有特别强调封装，这点并不像Java.所以提醒学过java的朋友，不用总是用java的语法特性来看待Golang,Golang本身对面向对象的特性做了简化的.

*/

import (
	"fmt"
	"go_code/go_study/16-面向对象三大特性/01-封装/model"
)

func main() {
	//创建一个account变量
	account := model.NewAccount("jzh11111", "999999", 40)
	if account != nil {
		fmt.Println("创建成功=", account)
	} else {
		fmt.Println("创建失败")
	}
}
