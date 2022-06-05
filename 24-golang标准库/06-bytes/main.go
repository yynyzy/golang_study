package main

import (
	"bytes"
	"fmt"
	"strings"
)

func test() {
	var i int = 100
	var b byte = 10
	b = byte(i)

	i = int(b)

	var s string = "hello world"
	b1 := []byte{1, 2, 3}

	s = string(b1)

	b1 = []byte(s)
	fmt.Println(s)
	fmt.Println(b1)
}

func test2() {
	s := "duoke360. com"
	b := []byte(s)
	b1 := []byte("duoke360")
	b2 := []byte("DuoKe360")
	strings.Contains("hello world", "hello")
	b3 := bytes.Contains(b, b1)
	fmt.Printf("b3: %v\n", b3)
	b3 = bytes.Contains(b, b2)
	fmt.Printf("b3: %v\n", b3)

}
func testCount() {
	s := []byte("he11ooooooooo")
	sep1 := []byte("h")
	sep2 := []byte("1")
	sep3 := []byte("o")
	fmt.Println(bytes.Count(s, sep1)) //1
	fmt.Println(bytes.Count(s, sep2)) //2
	fmt.Println(bytes.Count(s, sep3)) //9

}
func testRepeat() {
	//Repeat
	b := []byte("hi")
	fmt.Println(string(bytes.Repeat(b, 1))) //hi
	fmt.Println(string(bytes.Repeat(b, 3))) //hihihi
}
func testReplace() {
	//Replace
	s := []byte("hel1o ,world")
	old := []byte("o")
	news := []byte("ee")
	fmt.Println(string(bytes.Replace(s, old, news, 0))) //hello,world
	fmt.Println(string(bytes.Replace(s, old, news, 1))) //hellee ,world
	fmt.Println(string(bytes.Replace(s, old, news, 2))) //hellee,weerld
	//-1表示有几次替换几次
	fmt.Println(string(bytes.Replace(s, old, news, -1))) //hellee ,weerld
}

func testRunes() {
	//Runes
	s := []byte("你好世界")
	r := bytes.Runes(s)
	fmt.Println("转换前字符串的长度: ", len(s)) //12
	fmt.Println("转换后字符串的长度: ", len(r)) //4
}

func testJoin() {
	//Join
	s2 := [][]byte{[]byte("你好"), []byte("世界")}
	sep4 := []byte(",")
	fmt.Println(string(bytes.Join(s2, sep4))) //你好,世界
	sep5 := []byte("#")
	fmt.Println(string(bytes.Join(s2, sep5))) //你好#世界
}
func main() {
	// test()
	// test2()
	// testCount()
	// testRepeat()
	// testReplace()
	// testRunes()
	testJoin()
}
