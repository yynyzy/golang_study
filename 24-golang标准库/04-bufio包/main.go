package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

//ReadString
func test1() {
	f, _ := os.Open("go_study\\24-golang标准库\\04-bufio包\\b.txt")
	defer f.Close()
	r2 := bufio.NewReader(f)

	for {
		s, err := r2.ReadString('\n')
		fmt.Println(s)
		if err == io.EOF {
			break
		}
	}
}

func test2() {
	s := strings.NewReader("ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	br := bufio.NewReader(s)
	p := make([]byte, 10)
	for {
		n, err := br.Read(p)
		if err == io.EOF {
			break
		} else {
			fmt.Printf("string(p):%v\n", string(p[:n]))
		}

	}
}

//ReadByte
func test3() {
	s := strings.NewReader("ABCDEFG")
	br := bufio.NewReader(s)
	c, _ := br.ReadByte()
	fmt.Printf("%c\n", c)
	c, _ = br.ReadByte()
	fmt.Printf("%c\n", c)

	br.UnreadByte() //回退一个字节
	c, _ = br.ReadByte()
	fmt.Printf("%c\n", c)
}

//ReadLine
func test4() {
	s := strings.NewReader("ABC \nDEF\r\nGHI\r\nGAC")
	br := bufio.NewReader(s)

	//isPrefix 是不是前缀（一行特别多，后面的三会省略）
	w, isPrefix, _ := br.ReadLine()
	fmt.Printf("%q %v\n", w, isPrefix) //"ABC " false

	w, isPrefix, _ = br.ReadLine()
	fmt.Printf("%q %v\n", w, isPrefix) //	"DEF" false
	w, isPrefix, _ = br.ReadLine()
	fmt.Printf("%q %v\n", w, isPrefix) //	"GHI" false
	w, isPrefix, _ = br.ReadLine()
	fmt.Printf("%q %v\n", w, isPrefix) //	"GAC" false

	w, isPrefix, _ = br.ReadLine()     //第五次就没行读了
	fmt.Printf("%q %v\n", w, isPrefix) //	"" false
}

//ReadSlice
func test5() {
	//ReadSlice读取直到第一次遇到delim字节，返回缓冲里的包含已读取的数据和delim字节的切片。该返回值只在下一次读取操作之前合法。如果Readslice放在在读取到delim之前遇到了错误，它会返回在错误之前读取的数据在缓冲中的切片以及该错误(一般是io.EOF)。如果在读取到delim之前缓冲就被写满了，ReadSlice失败并返回ErrBufferFull。

	s := strings.NewReader("ABC,DEF,GHI,JKL")
	br := bufio.NewReader(s)
	w, _ := br.ReadSlice(',')
	fmt.Printf("%q\n", w)
	w, _ = br.ReadSlice(',')
	fmt.Printf("%q\n", w)
	w, _ = br.ReadSlice(',')
	fmt.Printf("%q\n", w)
	w, _ = br.ReadSlice(',')
	fmt.Printf("%q\n", w)
}

//ReadBytes
func test6() {
	s := strings.NewReader("ABC DEF GHI JKL ")
	br := bufio.NewReader(s)
	w, _ := br.ReadBytes(' ') //以空格为 delim
	fmt.Printf("%qin", w)
	w, _ = br.ReadBytes(' ')
	fmt.Printf("%q\n", w)
	w, _ = br.ReadBytes(' ')
	fmt.Printf("%q\n", w)
}

//WriteTo
func test7() {
	s := strings.NewReader("ABCEFGHIJKLMN")
	br := bufio.NewReader(s)
	// b := bytes.NewBuffer(make([]byte, 0))

	//open 打开文件只能只读，openfile 可写可读
	f, _ := os.OpenFile("go_study\\24-golang标准库\\04-bufio包\\c.txt", os.O_RDWR, 0777)
	defer f.Close()
	br.WriteTo(f)
	// fmt.Printf("%s\n", b)
}

//带缓冲区的写
func test8() {
	f, _ := os.OpenFile("go_study\\24-golang标准库\\04-bufio包\\c.txt", os.O_RDWR, 0777)
	defer f.Close()
	w := bufio.NewWriter(f)
	w.WriteString("hello YZY")
	w.Flush() //将缓冲区数据真正写入，（否则还在缓冲区）
}

//Available() 和 Buffered()
func test9() {
	b := bytes.NewBuffer(make([]byte, 0))
	bw := bufio.NewWriter(b)
	fmt.Println(bw.Available()) // 4096  ，缓冲区剩余多少字节容量
	fmt.Println(bw.Buffered())  // 0	，已缓冲的字节数

	bw.WriteString("ABCDEFGHIKLMN")
	fmt.Println(bw.Available()) // 4083
	fmt.Println(bw.Buffered())  // 13
	fmt.Printf("%q\n", b)       // ""

	bw.Flush()
	fmt.Println(bw.Available()) //4096
	fmt.Println(bw.Buffered())  //0
	fmt.Printf("%q\n", b)       //"ABCDEFGHIKLMN"

}
func main() {
	// test1()
	// test2()
	// test3()
	// test4()
	// test5()
	// test6()
	// test7()
	// test8()
	test9()
}
