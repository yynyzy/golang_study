package main

import "fmt"

func main() {
	type ValNode struct {
		row int
		col int
		val int
	}

	//原始数组
	var chessMap = [6][7]int{}
	chessMap[5][3] = 11
	chessMap[2][6] = 8

	//3.转成稀疏数组。->算法
	//思路
	//(1).遍历chessMap, 如果我们发现有一个元素的值不为0，创建一个node结构体
	//(2).将其放入到对应的切片即可

	var SparseArr = []ValNode{}
	//标准的一个稀疏数组应该还有一个记录元素的一维数组的规模(行和列，默认值)
	initValNode := ValNode{
		row: 6,
		col: 7,
		val: 0,
	}
	SparseArr = append(SparseArr, initValNode)

	for i, _ := range chessMap {
		for j, v := range chessMap[i] {
			if v != 0 {
				node := ValNode{
					row: i,
					col: j,
					val: v,
				}
				SparseArr = append(SparseArr, node)
			}
		}
	}

	//输出稀疏数组
	fmt.Println("当前稀疏数组是:...")
	for i, valNode := range SparseArr {
		fmt.Printf("%d: %d %d %d\n", i, valNode.row, valNode.col, valNode.val)
	}
}
