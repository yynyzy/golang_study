package main

import (
	"bufio"
	"fmt"
	"strings"
)

/*
Scanner类型提供了方便的读取数据的接口，如从换行符分隔的文本里读取每一行。成功调用的Scan方法会逐步提供文件的token,跳过token之间的字节。token由 SplitFunc 类型的分割函数指定; 默认的分割函数会将输入分割为多个行，并去掉行尾的换行标志。本包预定义的分割函数可以将文件分割为行、字节、unicode码值空白分隔的word.调用者可以定制自己的分割函数。扫描会在抵达输入流结尾、遇到的第一个 IO错误、token过大不能保存进缓冲时，不可恢复的停止。当扫描停止后，当前读取位置可能会远在最后一个获得的token后面。需要更多对错误管理的控制或token很大，或必须从reader连续扫描的程序，应使用bufio.Reader代替。
*/

func test_Split() {
	//func (s *Scanner) Split(split SplitFunc)
	//Split设置该Scanner的分割函数。本方法必须在Scan之前调用。
	s := strings.NewReader("ABC DEF GHI JKL")
	bs := bufio.NewScanner(s)
	bs.Split(bufio.ScanWords) //以空格作为分隔符，（ ScanBytes 以字节作为分隔符），（ ScanRunes 可以显示中文不乱码）
	for bs.Scan() {
		fmt.Println(bs.Text())
	}
}
func main() {
	test_Split()
}
