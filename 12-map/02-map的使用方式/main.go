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

	//第三种方式(也可以继续增加)
	heroes := map[string]string{
		"hero1": "yyn",
		"hero2": "yzy",
	}
	fmt.Println("heroes=", heroes)

	/*
			案例
		课堂练习:演示一个key-value的value是map的案例
		比如:我们要存放3个学生信息，每个学生有 name和 sex信息
		思路:map[string]map[string]string
		**/
	studentMap := make(map[string]map[string]string)

	studentMap["stu01"] = make(map[string]string, 3)
	studentMap["stu01"]["name"] = "tom"
	studentMap["stu01"]["sex"] = "男"
	studentMap["stu01"]["address"] = "北京长安街~"

	studentMap["stu02"] = make(map[string]string, 3)
	studentMap["stu02"]["name"] = "Marry"
	studentMap["stu02"]["sex"] = "女"
	studentMap["stu02"]["address"] = "上海黄浦江~"

	fmt.Println(studentMap)
	fmt.Println(studentMap["stu02"]["name"])

}
