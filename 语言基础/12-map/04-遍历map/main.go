package main

import "fmt"

func main() {
	//使用for-range遍历map
	//第二种 方式
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "天津"
	cities["no3"] = "上海"
	for key, val := range cities {
		fmt.Printf("k=%v v=%v\n", key, val)
	}
	fmt.Println("cities有", len(cities), "对key-value")

}
