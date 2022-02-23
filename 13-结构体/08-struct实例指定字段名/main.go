package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func main() {
	// 方式一
	var stu1 Student = Student{"小明", 18}
	stu2 := Student{"小红", 19}

	//方式二
	var stu3 Student = Student{
		Name: "小严",
		Age:  20,
	}
	stu4 := Student{
		Name: "小远",
		Age:  19,
	}
	fmt.Println(stu1, stu2, stu3, stu4)

	//方式2，返回结构体的指针类型(!!!)
	var stu5 *Student = &Student{"小王", 29} //stu5-->地址---》结构体数据[xxxx,xxx]
	stu6 := &Student{"小王~", 39}

	//在创建结构体指针变量时，把字段名和字段值写在一起，这种写法，就不依赖字段的定义顺序.
	var stu7 = &Student{
		Name: "小李", Age: 49,
	}
	stu8 := &Student{
		Age: 59, Name: "小李~",
	}
	fmt.Println(*stu5, *stu6, *stu7, *stu8)

}
