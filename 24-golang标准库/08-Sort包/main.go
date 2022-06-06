package main

import (
	"fmt"
	"sort"
)

func test_Sort() {
	var s []int = []int{2, 1, 5, 4, 3}
	sort.Ints(s) //顺序排序
	fmt.Println("s=", s)
}
func main() {

	test_Sort()

}
