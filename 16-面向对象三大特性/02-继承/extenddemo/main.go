package main

import "fmt"

/*
继承的深入讨论
	1)结构体可以使用嵌套匿名结构体所有的字段和方法，即:首字母大写或者小写的字段、方法，都可以使用。
	2)匿名结构体字段访问可以简化
	3) 当结构体和匿名结构体有相同的字段或者方法时，编译器采用就近访问原则访问，如希望访问匿名结构体的字段和方法，可以通过匿名结构体名来
*/

//1.小学生
type Student struct {
	Name  string
	Age   int
	Score int
}

//显示他的成绩
func (stu *Student) ShowInfo() {
	fmt.Printf("学生名=%v年龄=%v成绩=%v\n", stu.Name, stu.Age, stu.Score)
}
func (stu *Student) SetScore(score int) {
	stu.Score = score
}

//1.小学生
type Pupil struct {
	Student //嵌入了 Student 的匿名结构体
}

func (p *Pupil) testing() {
	fmt.Println("小学生正在考试.....")
}

//1.大学生
type Graduate struct {
	Student //嵌入了 Student 的匿名结构体
}

func (p *Graduate) testing() {
	fmt.Println("大学生正在考试.....")
}

func main() {
	//当我们对结构体嵌入了匿名结构体使用方法会发生变化
	pupil := &Pupil{}
	pupil.Student.Name = "tom~"
	pupil.Student.Age = 8
	pupil.testing()
	pupil.Student.SetScore(70)
	pupil.Student.ShowInfo()
	graduate := &Graduate{}
	graduate.Student.Name = "mary~"
	graduate.Student.Age = 28
	graduate.testing()
	graduate.Student.SetScore(99)
	graduate.Student.ShowInfo()

}
