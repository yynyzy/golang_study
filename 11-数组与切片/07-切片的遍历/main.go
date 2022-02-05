package main

import "fmt"

func main() {
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	slice := arr[1:3]

	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%v] =%v\n", i, slice[i])
	}
	for i, v := range slice {
		fmt.Printf("slice[%v] =%v\n", i, v)
	}
}
