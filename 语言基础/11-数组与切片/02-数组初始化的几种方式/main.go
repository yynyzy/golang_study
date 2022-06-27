package main

import (
	"fmt"
)

func main() {
	var numsArray01 [3]int = [3]int{1, 2, 3}
	var numsArray02 = [3]int{1, 2, 3}
	var numsArray03 = [...]int{6, 7, 8} //可以指定元素值对应的下标...
	var names = [3]string{1: "tom", 0: "jack", 2: "marry"}

	fmt.Println(numsArray01)
	fmt.Println(numsArray02)
	fmt.Println(numsArray03)
	fmt.Println(names)
}
