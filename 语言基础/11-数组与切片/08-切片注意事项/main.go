package main

import "fmt"

/*
·切片初始化时 var slice = arr[startlndex:endIndex]
 说明:从arr数组下标为startlndex，取到下标为endIndex的元素(不含arr[endIndex])。

·切片初始化时，仍然不能越界。范围在[O-len(arr)]之间，但是可以动态增长
	1.var slice = arr[0:end]可 以简写var slice = arr[:end]
	2.var slice = arr[start:len(arr)]可以简写: var slice = arr[start:]
	3.var slice = arr[O:len(arr)]可以简写: var slice = arr[:]

·cap是一个内置函数，用于统计切片的容量，即最大可以存放多少个元素。
·切片定义完后，还不能使用，因为本身是一个空的，需要让其引用到一个数组，或者make一个空间供切片来使用
·切片可以继续切片
**/

func main() {
	//用append内置函数，可以对切片进行动态追加

	/*1.
	   切片append操作的底层原理分析:
	   1)切片append操作的本质就是对数组扩容
	   2) go底层会创建一下新的数组newArr(如果没有超过cap时，扩容直接在尾部追加，当新数组长度超过了cap时才会拥有新的地址)
	   3)将slice原来包含的元素拷贝到新的数组newArr（或原数组）
	   4) slice重新引用到newArr（或原数组）
	   5)注意newArr是在底层来维护的，程序员不可见.
	**/

	var slice []int = []int{1, 2, 3, 4}
	fmt.Printf("第一次slice:%v,slice的地址=%p\n", slice, &slice)
	slice = append(slice, 100, 200, 300)
	fmt.Printf("第二次slice:%v,slice的地址=%p\n", slice, &slice)

	//可以将切片slice2追加给slice2
	slice = append(slice, slice...)
	fmt.Printf("第三次slice:%v,slice的地址=%p\n", slice, &slice)

	/*2.
		内置copy(par1,par2)函数
		·append会扩容，copy不会扩容
		·两个参数都必须是切片
		·将第二个切片复制到第一个切片
	**/
	var slice3 = []int{1, 2, 3, 4, 5}
	slice4 := make([]int, 10)
	copy(slice4, slice3)
	fmt.Println("slice3", slice3)
	fmt.Println("slice4", slice4)
	slice3[1] = 100 //数据空间独立，不会互相影响
	fmt.Println("slice3", slice3)
	fmt.Println("slice4", slice4)

}
