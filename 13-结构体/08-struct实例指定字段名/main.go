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
}
