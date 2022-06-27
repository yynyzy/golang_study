package main

//定义emp
type Emp struct {
	Id   int
	Name string
	Next *Emp
}

//定义EmpLink
//我们这里的EmpLink不带表头,即第一个结点就存放雇员
type EmpLink struct {
	Head *Emp
}

//方法待定..
//定义hashtable ,含有一个链表数组
type HashTable struct {
	LinkArr [7]EmpLink
}

//1。添加员工的方法,保证添加时,编号从小到大
func (this *EmpLink) Insert(emp *Emp) {
	cur := this.Head   //这是辅助指针
	var pre *Emp = nil //这是一个辅助指针pre在cur前面
	//如果当前的EmpLink就是一个空链表
	if cur == nil {
		this.Head = emp //完成
		return
	}
	//如果不是一个空链表,给emp找到对应的位置并插入
	//思路是让cur 和 emp比较,然后让pre保持在cur前面
	for {
		if cur != nil {
			//比较
			if cur.Id > emp.Id {
				//找到位置
				break
			}
			pre = cur
			cur = cur.Next
		} else {
			break
		}
	}
	pre.Next = emp
	emp.Next = cur
}

//给HashTable编写Insert雇员的方法.
func (this *HashTable) Insert(emp *Emp) {
	//使用散列函数,确定将该雇员添加到哪个链表
	linkNo := this.HashFun(emp.Id)
	//使用对应的链表添加
	this.LinkArr[linkNo].Insert(emp)
}

//编写一个散列方法
func (this *HashTable) HashFun(id int) int {
	return id % 7 //得到一个值,就是对于的链表的下标
}
func main() {

}
