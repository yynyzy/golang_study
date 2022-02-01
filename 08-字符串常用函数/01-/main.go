package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//1.统计字符串的长度，按字节 len(str)
	str := "hello北"
	//golang的编码统一为utf-8（ascil的字符（字母和数字）占一个字符，汉字占用3个字节）
	fmt.Println("str的长度=", len(str))

	//2.字符串遍历，同时处理有中文的问题 []rune(str)
	str2 := "hello北京"
	r := []rune(str2)
	for i := 0; i < len(r); i++ {
		fmt.Printf("字符=%c\n", r[i])
	}

	//3.字符串转整数：
	// n, err := strconv.Atoi("hello") //invalid syntax
	n, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("错误转换", err)
	} else {
		fmt.Println("转换的结果是", n)
	}

	//4.整数转字符串：
	str = strconv.Itoa(121)
	fmt.Printf("str=%v,str=%T\n", str, str)

	//5) 字符串转[]byte: var bytes = []byte("hello go")
	var bytes = []byte("he1lo go")
	fmt.Printf("bytes=%v\n", bytes)

	//6) []byte转字符串:str = string([]byte[97,98,99})
	str = string([]byte{97, 98, 99})
	fmt.Printf("str=%v\n", str)

	//7) 10进制转2，8，16进制:str = strconv.FormatInt(123，2),返回对应的字符串
	str = strconv.FormatInt(123, 2)
	fmt.Printf("123对应的二进制是=%v\n", str)
	str = strconv.FormatInt(123, 16)
	fmt.Printf("123对应的十六进制是=%v\n", str)

	//8)查找子串是否在指定的字符串中: strings.Contains("seafood","foo"') //true
	b := strings.Contains("seafood", "mary")
	fmt.Printf("b=%v\n", b)

	//9)统计一个字符串有几个指定的子串:strings.count(" ceheese","e")//4
	num := strings.Count("ceheese", "se")
	fmt.Printf("num=%v\n", num)

	//10)不区分大小写的字符串比较(==是区分字母大小写的):
	b = strings.EqualFold("abc", "Abc")
	fmt.Printf("b=%v\n", b)           //true
	fmt.Println("结果", "abc" == "Abc") // false //区分字母大小写

	//11)返回子串在字符串第一次出现的index值,如果没有返回-1:
	//strings .Index("NLT_abc", "abc") //4
	index := strings.Index("NLT_abcabcabc", "abc") //4
	fmt.Printf("index=%v\n", index)

	//12)返回子串在字符串最后一次出现的index
	//如没有返回-1 :  strings.LastIndex( "go golang", "go")
	index = strings.LastIndex("go golang", "go") //3
	fmt.Printf("index=%v\n", index)

	//将指定的子串替换成另外一个子串: strings.Replace("to go hello", "go","go语言",n)
	//n可以指定你希望替换几个,如果n=-1表示全部替换
	str2 = "go go hello"
	str = strings.Replace(str2, "go", "北京", -1) //返回一个新字符串
	fmt.Printf("str=%v\n", str)

	//按照指定的某个字符，为分割标识，将一个字符串拆分成字符串数组:
	//strings.split(""hello,wrold,ok", ",")
	strArr := strings.Split("hello,wrold,ok", ",	")
	fmt.Printf("strArr=%v\n", strArr)

}
