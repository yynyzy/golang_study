package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func test() {
	fmt.Printf("float64的最大值是:%.f\n", math.MaxFloat64)
	fmt.Printf(" float64的最小值是:%.f\n", math.SmallestNonzeroFloat64)
	fmt.Printf("float32的最大值是:%.f\n", math.MaxFloat32)
	fmt.Printf("float32的最小值是:%.f\n", math.SmallestNonzeroFloat32)
	fmt.Printf("Int8的最大值是:%d\n", math.MaxInt8)
	fmt.Printf("Int8的最小值是:%d\n", math.MinInt8)
	fmt.Printf("Uint8的最大值是:%d\n", math.MaxUint8)
	fmt.Printf("Int16的最大值是:%d\n", math.MaxInt16)
	fmt.Printf("Int16的最小值是:%d\n", math.MinInt16)
	fmt.Printf("Uint16的最大值是:%d\n", math.MaxUint16)
	fmt.Printf("Int32的最大值是:%d\n", math.MaxInt32)
	fmt.Printf("Int32的最小值是:%d\n", math.MinInt32)
	fmt.Printf("U1nt32的最大值是:%d\n", math.MaxUint32)
	fmt.Printf("Int64的最大值是:%d\n", math.MaxInt64)
	fmt.Printf(" Int64的最小值是:%d\n", math.MinInt64)
	fmt.Printf("圆周率默认为:%.200f\n", math.Pi)

}

func test2() { //分别返回整数和小数部分
	/*
		分别取整数和小数部分,函数签名如下:
			func Modf(f float64) (int float64, frac float64)
	*/
	Integer, Decimal := math.Modf(3.14159265358979)
	fmt.Printf("[3.14159265358979]的整数部分为:[%.f],小数部分为:[%.14f]\n", Integer, Decimal)
}
func init() {
	rand.Seed(time.Now().UnixNano())
}
func main() {
	// test()
	// test2()
	a := rand.Int()
	fmt.Printf("rand.Int:%d\n", a)
}
