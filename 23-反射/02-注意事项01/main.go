package main

/*
1) reflect.Value.Kind, 获取变量的类别，返回的是一个常量
2) Type是类型, Kind是类别，Type和Kind可能是相同的，也可能是不同的
比如：var num int= 10 num 的 Type 是 int , Kind也是 int
比如：var stu Student stu的 Type 是 pkg1.Student , Kind是 struct
*/
