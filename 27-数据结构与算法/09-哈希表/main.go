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

func main() {

}
