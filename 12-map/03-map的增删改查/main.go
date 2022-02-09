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

	//删除操作
	delete(cities, "no1")
	fmt.Println(cities)
	//当delete指定的key不存在时,删除不会操作,也不会报错
	delete(cities, "no4")
	fmt.Println(cities)

}
