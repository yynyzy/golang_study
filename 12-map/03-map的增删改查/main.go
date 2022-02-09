package main

import "fmt"

func main() {
	//增加和修改
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "天津"
	cities["no3"] = "北海"
	fmt.Println(cities)
	//因为no3这个key已经存在,因此下面的这句话就是修改
	cities["no3"] = "上海~"
	fmt.Println(cities)

	/*
		删除操作
			➢细节说明
			1)如果我们要删除map的所有key ,没有一个专门的方法一次删除， 可以遍历一下key,逐个删除
			2)或者map = make(...),. make一个新的，让原来的成为垃圾，被 gc 回收

	*/
	delete(cities, "no1")
	fmt.Println(cities)
	//当delete指定的key不存在时,删除不会操作,也不会报错
	delete(cities, "no4")
	fmt.Println(cities)

}
