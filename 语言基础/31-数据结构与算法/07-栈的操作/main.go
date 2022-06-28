package main

import (
	"errors"
	"fmt"
)

//使用数组来模拟一个栈的使用
type stack struct {
	MaxTop int    //表示我们栈最大可以存放数个数
	Top    int    //表示栈顶
	arr    [5]int //数组模拟栈
}

//入栈
func (this *stack) Push(val int) (err error) {
	//先判断栈是否满了
	if this.Top == this.MaxTop-1 {
		fmt.Println("stack full")
		return errors.New("stack full")
	}
	this.Top++ //放入数据
	this.arr[this.Top] = val
	return

}

//出栈
//入栈
func (this *stack) Pop() (val int, err error) {
	//先判断栈是否满了
	if this.Top == -1 {
		fmt.Println("stack Empty")
		return 0, errors.New("stack Empty")
	}
	val = this.arr[this.Top]
	this.Top-- //放入数据
	return val, nil

}

//遍历栈,注意需要从栈顶开始遍历
func (this *stack) List() {
	//先判断栈是否为空
	if this.Top == -1 {
		fmt.Println("stack empty")
		return
	}
	fmt.Println("栈的情况如下:")
	for i := this.Top; i >= 0; i-- {
		fmt.Printf("arr[%d]=%d\n", i, this.arr[i])
	}
}

func main() {
	stack := &stack{
		MaxTop: 5,
		Top:    -1,
	}
	//入栈
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	//显示
	stack.List()

	val, _ := stack.Pop()
	fmt.Println("出栈val=", val) //出栈
	//显示
	stack.List()

}
