package main

import "fmt"

func quickSort(arr []int) []int {
	var lens = len(arr)
	if lens <= 1 {
		return arr
	}
	var privot = arr[0]
	var middleArr = []int{privot}
	var leftArr = make([]int, 0, 0)
	var rightArr = make([]int, 0, 0)
	for i := 1; i < lens; i++ {
		if arr[i] < privot {
			leftArr = append(leftArr, arr[i])
		} else {
			rightArr = append(rightArr, arr[i])
		}
	}
	leftArr, rightArr = quickSort(leftArr), quickSort(rightArr)
	res := append(append(leftArr, middleArr...), rightArr...)
	return res
}
func main() {
	arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
	quickSort(arr)
	fmt.Println("quickSort(arr):", quickSort(arr))
}
