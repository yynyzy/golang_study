package main

import (
	"fmt"
)

/*
switch细节讨论
1) case后是一个表达式(即:常量值、变量、一个有返回值的函数等都可以)
2) case后的各个表达式的值的数据类型，必须和 switch的表达式数据类型一致
3) case后面可以带多个表达式，使用逗号间隔。比如case表达式1,表达式2..
4) case后面的表达式如果是常量值(字面量)，则要求不能重复
5) case后面不需要带break ,程序匹配到一个case后就会执行对应的代码块，然后退出switch，如果一个都匹配不到，则执行default
6)default语句不是必须的
7) switch后也可以不带表达式，类似if --else分支来使用。
8)switch后也可以直接声明/定义一个变量，分号结束，不推荐。
9）switch穿透-fallthrough，如果在case语句块后增加 fallthrough ,则会继续执行下一个case（只会穿透一层）,也叫switch穿透。

*/
func main() {
	/*
		switch分支结构
		switch 表达式{
			case 表达式1，表达式2，...：语句块1

			case 表达式1，表达式2，...：语句块1
			//这里可以有多个case

			default：
			语句块
		}
	*/

	var key byte
	fmt.Println("请输入字符a或b")
	fmt.Scanf("%c", &key)
	switch key {
	case 'a':
		fmt.Println("星期一")
	case 'b':
		fmt.Println("星期二")
	default:
		fmt.Println("星期天")
	}
}
