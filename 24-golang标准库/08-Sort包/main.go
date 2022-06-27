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

func test_SortString() {
	//字符串排序,现比较高位,相同的再比较低位
	//[]string
	ls := sort.StringSlice{"100", "42", "41", "3 ", "2"}
	fmt.Println(ls) //[ 100 42 41 3 2]
	sort.Strings(ls)
	fmt.Println(ls) //[188 2 3 41 42]

}

func test_SortFloat() {
	f := []float64{1.1, 4.4, 5.5, 3.3, 2.2}
	sort.Float64s(f)
	fmt.Printf("f: %v\n", f) //f=[1.1, 2.2 ,3.3 ,4.4, 5.5]

}

//汉字排序，依次比较byte大小
func test_SortString2() {
	ls := sort.StringSlice{"啊", "博", "次", "得", "饿", "周"}

	fmt.Println(ls) //[啊博次得饿周]
	sort.Strings(ls)
	for _, v := range ls {
		fmt.Println(v, []byte(v))
	}
}

//自定义类型 NewInts 实现 sort接口
type NewInts []uint

func (n NewInts) Len() int {
	return len(n)
}
func (n NewInts) Less(i, j int) bool {
	fmt.Println(i, j, n[i] < n[j], n)
	return n[i] < n[j]
}
func (n NewInts) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

// [][]int 类型排序
type testslice [][]int

func (l testslice) Len() int           { return len(l) }
func (l testslice) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l testslice) Less(i, j int) bool { return l[i][1] < l[j][1] } //比较第二位数

func main() {

	// test_Sort()
	// test_SortFloat()
	// test_SortString()
	// test_SortString2()

	// n := []uint{1, 3, 2}
	// sort.Sort(NewInts(n)) //将uint类型强转为 NewInts 类型
	// fmt.Println(n)

	ls := testslice{{1, 4}, {9, 3}, {7, 5}}
	fmt.Println(ls) //[[1 4][9 3][7 5]]
	sort.Sort(ls)
	fmt.Println(ls) //[[9 3][1 4][7 5]]
}
