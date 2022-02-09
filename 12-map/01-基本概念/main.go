package main

import "fmt"

/*
#基本语法：
	var map 变量名  map[keytype]valuetype

## key 可以是什么类型
golang中的map，的 key可以是很多种类型，比如 bool，数字，string，指针，channel ，还可以是只包含前面几个类型的接口，结构体,数组。通常为int 、 string
注意: slice, map还有 function不可以，因为这几个没法用 ==来判断

## valuetype 可以是什么类型
valuetype的类型和key基本一样，这里我就不再赘述了通常为:数字(整数,浮点数),string,map,struct

map声明的举例:
var a map[string]string
var a map[string]int
var a map[int]string
var a map[string]map[string]string
注意:声明是不会分配内存的，初始化需要make，分配内存后才能赋值和使用。

**/
func main() {
	var a map[string]string = make(map[string]string, 5)
	a["no1"] = "严以诺"
	a["no2"] = "严zy"
	a["no1"] = "严ss"
	a["no3"] = "严zy"
	fmt.Println(a)
	/*
		对上面代码的说明
			1) map在使用前一定要make
			2) map的 key是不能重复，如果重复了，则以最后这个key-value为准
			3) map的 value是可以相同的.
			4) map的key-value是无序
	*/
}
