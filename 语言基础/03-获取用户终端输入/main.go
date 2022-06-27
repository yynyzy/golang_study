package main

import (
	"fmt"
)

//在控制台接受用户输入信息
func main() {
	//方式一：使用 fmt.Scanln
	var name string
	var age byte
	var salary float64
	var CET6 bool
	// fmt.Println("请输入你的名字？")
	// fmt.Scanln(&name)
	// fmt.Println("请输入你的年龄？")
	// fmt.Scanln(&age)
	// fmt.Println("请输入你的薪水？")
	// fmt.Scanln(&salary)
	// fmt.Println("是否过英语6级？")
	// fmt.Scanln(&CET6)

	// fmt.Printf("你的名字%v，你的年龄%v，你的薪水%v，是否过六级%v", name, age, salary, CET6)

	//方式二：使用 fmt.Scanf
	fmt.Println("请输入你的名字,年龄,薪水,是否过英语6级？")
	fmt.Scanf("%v %d %f %b", &name, &age, &salary, &CET6)

	fmt.Printf("你的名字%s，你的年龄%v，你的薪水%v，是否过六级%v,请用空格分开", name, age, salary, CET6)
}
