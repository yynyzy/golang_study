package main

func main() {
	/*
		1)map是引用类型，遵守引用类型传递的机制，在一个函数接收map，修改后，会直接修改原来的map
		2)map的容量达到后，再想map增加元素，会自动扩容，并不会发生panic，也就是说map能动态的增长键(key-value)(！！！切片map不行哦)
		3)map的value也经常使用struct类型，更适合管理复杂的数据(比前面value是一个map更好)结构体
	*/
}
