package main

import "fmt"

//选择排序，从大到小排
func selectSort(arr *[5]int) {

	for i := 0; i < len(arr); i++ {
		var max = arr[i]
		var maxIndex = i

		for j := i + 1; j < len(arr); j++ {
			if max < arr[j] {
				max = arr[j]
				maxIndex = j
			}
		}
		if maxIndex != i {
			arr[i], arr[maxIndex] = arr[maxIndex], arr[i]
		}
	}
}
func main() {
	var arr = [5]int{10, 200, 80, 30, 5}
	selectSort(&arr)
	fmt.Println("选择排序后arr=", arr)
}
