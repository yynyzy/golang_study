package main

import "fmt"

func main() {
	var a map[string]string = make(map[string]string, 5)
	a["no1"] = "宋江"
	a["no2"] = "吴用"
	fmt.Println(a)

	//第二种方式(map容量自动扩增)
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "天津"
	cities["no3"] = "上海"
	fmt.Println(cities)

	//第三种方式
	heroes := map[string]string{
		"hero1": "yyn",
		"hero2": "yzy",
	}
	fmt.Println("heroes=", heroes)

}
