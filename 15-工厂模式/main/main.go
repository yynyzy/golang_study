package main

import (
	"fmt"
	"golang_study/15-工厂模式/model"
)

func main() {
	// stu := model.student{
	// 	Name:  "严以诺",
	// 	Score: 21,
	// }

	stu := model.NewStudent("yzy", 18)
	//student结构体首字母是小写，可以通过工厂模式来解决
	fmt.Println(*stu)
	fmt.Println("stu的score为=", stu.GetScore())
}
