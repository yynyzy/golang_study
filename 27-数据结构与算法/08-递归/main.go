package main

import "fmt"

func main() {

	//先创建一个二维数组，模拟迷宫
	//规则
	//1.如果元素的值为缸，就是墙
	//2．如果元素的值为e，是没有走过的点
	//3．如果元素的值为2，是一个通路
	//4．如果元素的值为3,是走过的点,但是走不通
	var myMap [8][7]int

	//先把地图的最上和最下设置为1
	for i := 0; i < 7; i++ {
		myMap[0][i] = 1
		myMap[7][i] = 1
	}
	//先把地图的最左和最右设置为1
	for i := 0; i < 8; i++ {
		myMap[i][0] = 1
		myMap[i][6] = 1
	}
	//输出地图
	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(myMap[i][j], " ")
		}
		fmt.Println()
	}
}
