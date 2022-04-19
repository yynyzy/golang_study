package main
import (
	fmt
)
//使用一个结构体管理队列
type Queue struct {
	maxSize int
	array   [4]int //数组=>模拟队列
	front   int    //表示指向队列首
	rear    int    //表示指向队列的尾部
}

//添加数据到队列
func (this *Queue) AddQueue(val int) (err error) {
	//先判断队列是否已满
	if this.rear == this.maxSize-1 { //重要重要的提示;rear是队列尾部(含最后元素
		return errors.New("queue full")
	}
	this.rear++ //rear后移
	this.array[this.rear] = val
	return
}

func (this *Queue) showQueue() {
	fmt.Println("队列当前的情况是:") //this.front不包含队首的元素
	for i := this.front + 1; i <= this.rear; i++ {
		fmt.Printf("array[%d]=%d\t", i, this.array[i])
	}
}

//编写一个主函数测试,测试
func main() {
	//先创建一个队列
	queue := &Queue{
		maxSize: 5,
		front:   -1,
		rear:    -1,
	}
	var key string
	for {
		fmt.Println("1.输入add表示添加数据到队列")
		fmt.Println("2.输入get表示从队列获取数据")
		fmt.Println("3.输入show表示显示队列")
		fmt.Println("4.输入exit表示显示队列")
		fmt.scanln(&key)
		switch key {
		case "add":
			fmt.Println("输入你要入队列数")
		}
	}
}
