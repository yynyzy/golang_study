# **Go切片与技巧**

# *go 遍历 chan 的3种方法*
```go
【一】for v:=range p{//遍历出来的数据，每一层创建一个容器v,v作用域是一层循环，一层结束后释放，下一层再建
		fmt.Println(v)//循环的条件是   p中有数据
	  }
 
【二】	for{//无条件循环，最后通过break跳出循环   注意循环不会造成栈溢出，因为一层结束后释放本层资源
			if v,ok:=<-p;ok{//每一层遍历都创建 v ok,同第一种方式
				fmt.Println(v)
			}else{
				break
			}
		}

【三】v,ok:=<-p
	 for ok{//循环条件是 ok为真 //循环出来的数据放在循环体外
		fmt.Println(v)
		v,ok=<-p
	 }
【三也可以这样写】
      v,ok ；= <-p
      for{
         if ok{
           fmt.Println(v)
		   v,ok=<-p
         }else{
           break
         }       
      }
【共同点：都需要循环！！！！。】
```

# Printf()、Sprintf()、Fprintf()函数的区别用法是什么？
简单来说，这三个 fmt.Print 函数的作用分别是：
*Printf*
格式化并输出到标准输出
*Sprintf*
格式化并返回一个字符串，不直接输出
*Fprintf*
格式化并输出到指定的输出流（文件、标准输出等）

# Golang 函数和 C 函数深度对比
## *一、C 语言函数深究*
我们准备一段简单的函数调用代码。
```c
#include <stdio.h>  
int func(int p){   
    return 1;
}  
int main()  
{  
    int i;  
    for(i=0; i<100000000; i++){  
        func(2);  
    }  
    return 0;  
}
```
用 gcc 来查看下汇编代码。
//# gcc -S main.c

汇编源码如下：
```
func:
        movl    %edi, -4(%rbp)
        movl    %esi, -8(%rbp)
        movl    %edx, -12(%rbp)
        movl    %ecx, -16(%rbp)
        movl    %r8d, -20(%rbp)
        movl    $1, %eax
        
main:
        movl    $5, %r8d
        movl    $4, %ecx
        movl    $3, %edx
        movl    $2, %esi
        movl    $1, %edi
        call    func
       
        movl    $0, %eax
```
可以看到，在C语言中：

主要通过寄存器传递参数
所以，C 语言函数的性能杠杠的。寄存器是整个计算机体系结构中访问最最快的存储了。只有当参数数量大于 6 的时候，才开始使用栈。

固定 eax 寄存器返回数据
因为固定使用 eax 寄存器做返回数据之用，所以在 C 语言中无法支持多个返回值。我们接下来看看 Golang 是如何支持多个返回值的。

## *二、Golang 函数深究*
同样先写一段最简单的函数调用代码。
```go
package main

func myFunction(p1, p2, p3,p4, p5 int) (int,int) {
 var a int = p1+p2+p3+p4+p5
 var b int = 3
 return a,b
}

func main() {
 myFunction(1, 2, 3, 4, 5)
}
```
然后查看其汇编代码。

//为了方便查看，使用-N -l 参数，能阻止编译器对汇编代码的优化
#go tool compile -S -N -l main.go > main.s
结果是这样的：
```
"".main STEXT size=95 args=0x0 locals=0x38
  0x000f 00015 (main.go:7)  SUBQ $56, SP  //在栈上分配56字节
  0x0013 00019 (main.go:7)  MOVQ BP, 48(SP) //保存BP
  0x0018 00024 (main.go:7)  LEAQ 48(SP), BP

        0x001d 00029 (main.go:8)        MOVQ    $1, (SP) //第一个参数入栈
        0x0025 00037 (main.go:8)        MOVQ    $2, 8(SP) //第二个参数入栈
        0x002e 00046 (main.go:8)        MOVQ    $3, 16(SP) //第三个参数入栈
        0x0037 00055 (main.go:8)        MOVQ    $4, 24(SP) //第四个参数入栈
        0x0040 00064 (main.go:8)        MOVQ    $5, 32(SP) //第五个参数入栈
        0x0049 00073 (main.go:8)        CALL    "".myFunction(SB)

"".myFunction STEXT nosplit size=99 args=0x38 locals=0x18
        0x000e 00014 (main.go:3) MOVQ $0, "".~r5+72(SP)
        0x0017 00023 (main.go:3) MOVQ $0, "".~r6+80(SP)
        0x0020 00032 (main.go:4) MOVQ "".p1+32(SP), AX
        0x0025 00037 (main.go:4) ADDQ "".p2+40(SP), AX
        0x002a 00042 (main.go:4) ADDQ "".p3+48(SP), AX
        0x002f 00047 (main.go:4) ADDQ "".p4+56(SP), AX
        0x0034 00052 (main.go:4) ADDQ "".p5+64(SP), AX
        0x004b 00075 (main.go:6) MOVQ AX, "".~r5+72(SP)
        0x0054 00084 (main.go:6) MOVQ AX, "".~r6+80(SP)
```
可以看到，在Golang中：

使用栈来传递参数
栈是位于内存之中的，虽然有 CPU 中 L1、L2、L3的帮助，但平均每次访问性能仍然和寄存器没法比。所以 Golang 的函数调用开销肯定会比 C 语言要高。后面我们将用一个实验来进行量化的比较。

使用栈来返回数据
不像 C 语言那样固定使用一个 eax 寄存器，Golang 是使用栈来返回值的。这就是为啥 Golang 可以返回多个值的根本原因。

最后，性能开销对比
我们的测试方法简单粗暴，直接调用空函数 1 亿次，再统计计算平均耗时。

C函数编译运行测试：
```c++
#include <stdio.h>  
int func(int p){   
    return 1;
}  
int main()  
{  
    int i;  
    for(i=0; i<100000000; i++){  
        func(2);  
    }  
    return 0;  
}
//# gcc main.c -o main
//# time ./main
```
第一次执行耗时大约是 0.339 s。

但这个耗时中包含了两块。一块是函数调用开销，另外一块是 for 循环的开销(其它的代码调用因为只有 1 次，而函数调用和 for 循环都有 1 亿次，所以直接就可以忽略了)。

所以我们得减去 for 循环的开销。接着我手工注释掉对函数的调用，只是空循环 100000000 次。
```go
int main()  
{  
    int i;  
    for(i=0; i<100000000; i++){  
    //    func(2);  
    }  
    return 0;  
}
```
这次总耗时是 0.314 s。

这样就计算出平均每次函数调用耗时 = (0.339s - 0.314s) / 100000000 = 0.25ns

Golang函数编译运行
```go
func hello(a int) int {
 return 2
}

func main(){
 for i:=0; i<100000000; i++ {
  hello(1)
 }
}
//# go build -gcflags="-m -l" main.go
```
同样采用上述方法测出平均每次函数调用耗时 = (0.302s - 0.056 s) / 100000000 = 2.46ns

可见 Golang 的函数调用性能还是比 C 要差一些。但再给大家个参考一下 PHP 的数据，之前我测过 PHP7 每次函数调用开销大约在 50 ns 左右。所以 Golang 虽然比不上 C，但总的来说性能还是不错的。


## ·切片的创建
*切片底层的数据结构*（src/runtime/slice.go）：
```go
type slice struct {
 array unsafe.Pointer // 指针指向底层数组
 len   int  // 切片长度
 cap   int  // 底层数组容量
}
```
切片的类型规范是[]T，其中T是切片元素的类型。与数组类型不同，切片类型没有指定长度。

### *创建切片有多种方式*：变量声明、切片字面量、make创建、new创建、从切片/数组截取。

*1、变量声明*
var s []byte //这种声明的切片变量，默认值是nil，容量和长度默认都是0。

*2、切片字面量*
当使用字面量来声明切片时，其语法与使用字面量声明数组非常相似：
```go
a := []string{"johnny", "太白技术"} // 这是切片声明，少了长度的指定
```

*3、make创建：*
除了上面创建方式外，使用内置函数make可以指定长度和容量来创建（具体函数：func make([]T, len, cap) []T）

其中 T 代表要创建的切片的元素类型。该make函数采用类型、长度和可选容量。调用时，make分配一个数组并返回一个引用该数组的切片。

首先看下空切片：
```go
 y := make([]int, 0)
 fmt.Println(cap(y))   // 0
 fmt.Println(len(y))   // 0
 fmt.Println(y == nil) // false
下面这个例子中创建了长度为5，容量等于5的切片

 var s []byte
 s = make([]byte, 5, 5) // 指定了长度和容量
 fmt.Println(s) // 输出：[0 0 0 0 0]

当容量参数被省略时，它默认为指定的长度：

s := make([]byte, 5)
fmt.Println(s) // 输出也是 [0 0 0 0 0]
内置len和cap函数检查切片的长度和容量：

 fmt.Println(len(s)) // 5
 fmt.Println(cap(s)) // 5
```

*4、使用new*
使用new创建的切片是个nil切片。

 n := *new([]int)
 fmt.Println(n == nil) // true

*5、从切片/数组截取*
切片可以基于数组和切片来创建。


## ·切片操作
内置函数append()用于向切片中追加元素。

 n := make([]int, 0)
 n = append(n, 1)                 // 添加一个元素
 n = append(n, 2, 3, 4)           // 添加多个元素
 n = append(n, []int{5, 6, 7}...) // 添加一个切片
 fmt.Println(n)                   // [1 2 3 4 5 6 7]
当append操作的时候，切片容量如果不够，会触发扩容，接着上面的例子：

 fmt.Println(cap(n)) // 容量等于8
 n = append(n, []int{8, 9, 10}...)
 fmt.Println(cap(n)) // 容量等于16，发生了扩容
当一开始容量是8，后面追加了切片[]int{8, 9, 10}之后，容量变成了16。关于详细的扩容机制，我们后面文章再聊。

## ·切片表达式
Go语言提供了两种切片的表达式：

1、简单表达式
2、扩展表达式

*1、简单的表达式*
简单切片表达式的格式[low:high]。

如下面例子，n为一个切片，当用这个表达式[1:4]表示的是左闭右开[low, high)区间截取一个新的切片（例子中结果是[2 3 4]），切片被截取之后，截取的长度是high-low。
```go
 n := []int{1, 2, 3, 4, 5, 6}
 fmt.Println(n[1:4])      // [2 3 4]
切片表达式的开始low和结束索引high是可选的；它们分别默认为零和切片的长度：

n := []int{1, 2, 3, 4, 5, 6}
fmt.Println(n[:4])      // [1 2 3 4]，:前面没有值，默认表示0
fmt.Println(n[1:])       // [2 3 4 5 6]，:后面没有值，默认表示切片的长度
边界问题
1、当n为数组或字符串表达式n[low:high]中low和high的取值关系：
0 <= low <=high <= len(n)

2、当n为切片的时候，表达式n[low:high]中high最大值变成了cap(n)，low和high的取值关系：
0 <= low <=high <= cap(n)
```

不满足以上条件会发送越界panic。
截取字符串
注意截取字符串，产生的也是新的字符串。
```go
 s := "hello taibai"
 s1 := s[6:]
 fmt.Println(s1)                 // taibai
 fmt.Println(reflect.TypeOf(s1)) // string
 ```

*扩展表达式*
简单表达式生产的新切片与原数组或切片会共享底层数组，虽然避免了copy，但是会带来一定的风险。下面这个例子当新的n1切片append添加元素的时候，覆盖了原来n的索引位置4的值，导致你的程序可能是非预期的，从而产生不良的后果。
```go
 n := []int{1, 2, 3, 4, 5, 6}
 n1 := n[1:4]
 fmt.Println(n)       // [1 2 3 4 5 6]
 fmt.Println(n1)      // [2 3 4]
 n1 = append(n1, 100) // 把n的索引位置4的值从原来的5变成了100
 fmt.Println(n)       // [1 2 3 4 100 6]
 ```
Go 1.2[3]中提供了一种可以限制新切片容量的表达式：

n[low:high:max]

max表示新生成切片的容量，新切片容量等于max-low，表达式中low、high、max关系：

0 <= low <= high <= max <= cap(n)

继续刚才的例子，当计算n[1:4]的容量，用cap得到值等于5，用扩展表达式n[1:4:5]，用cap重新计算得到新的容量值（5-1）等于4：

 fmt.Println(cap(n[1:4])) // 5
 fmt.Println(cap(n[1:4:5])) // 4
关于容量
n[1:4]的长度是3好理解（4-1），容量为什么是5？

因为切片n[1:4]和切片n是共享底层空间，所以它的容量并不等于他的长度3，根据1等于索引1的位置（等于值2），从值2这个元素开始到末尾元素6，共5个，所以n[1:4]容量是5。

如果append超过切片的长度会重新生产一个全新的切片，不会覆盖原来的：
```go
 n2 := n[1:4:5]         // 长度等于3，容量等于4
 fmt.Printf("%p\n", n2) // 0xc0000ac068
 n2 = append(n2, 5)
 fmt.Printf("%p\n", n2) // 0xc0000ac068
 n2 = append(n2, 6)
 fmt.Printf("%p\n", n2) // 地址发生改变，0xc0000b8000
```

## ·切片技巧
### *AppendVector*
这个是添加一个切片的操作，上面我们在切片操作中已经介绍过。
```go
a = append(a, b...)
```

### *Copy*
这边给我们展示了三种copy的写法：
```go
b := make([]T, len(a))
copy(b, a)

// 效率一般比上面的写法慢，但是如果有更多，他们效率更好
b = append([]T(nil), a...)
b = append(a[:0:0], a...)

// 这个实现等价于make+copy。
// 但在Go工具链v1.16上实际上会比较慢。
b = append(make([]T, 0, len(a)), a...)
```

### *Cut*
截掉切片[i,j）之间的元素：
```go
a = append(a[:i], a[j:]...)
```

### *Cut（GC）*
上面的Cut如果元素是指针的话，会存在内存泄露，所以我们要对删除的元素设置nil，等待GC。
```go
copy(a[i:], a[j:])
for k, n := len(a)-j+i, len(a); k < n; k++ {
 a[k] = nil // or the zero value of T
}
a = a[:len(a)-j+i]
```

### *Delete*
删除索引位置i的元素：
```go
a = append(a[:i], a[i+1:]...)
// or
a = a[:i+copy(a[i:], a[i+1:])]
```

### *Delete（GC）*
删除索引位置i的元素：
```go
copy(a[i:], a[i+1:])
a[len(a)-1] = nil // or the zero value of T
a = a[:len(a)-1]
```

### *Delete without preserving order*
删除索引位置i的元素，把最后一位放到索引位置i上，然后把最后一位元素删除。这种方式底层并没有发生复制操作。
```go
a[i] = a[len(a)-1] 
a = a[:len(a)-1]
```

### *Delete without preserving order（GC）*
上面的删除操作，元素是一个指针的类型或结构体指针字段，会存在最后一个元素不能被GC掉，造成泄露，把末尾的元素设置nil，等待GC。
```go
a[i] = a[len(a)-1]
a[len(a)-1] = nil
a = a[:len(a)-1]
```

### *Expand*
这个本质上是多个append的组合操作。
```go
a = append(a[:i], append(make([]T, j), a[i:]...)...)
```


### *Extend*
用新列表扩展原来的列表
```go
a = append(a, make([]T, j)...)
```

### *Filter (in place)*
下面代码演示原地删除Go切片元素：
```go
n := 0
for _, x := range a {
 if keep(x) {
  a[n] = x
  n++
 }
}
a = a[:n]
```

### *Insert*
```go
a = append(a[:i], append([]T{x}, a[i:]...)...)
第二个append会产生新的切片，产生一次copy，可以用以下代码方式，可免去第二次的copy：

s = append(s, 0 /* 先添加一个0值*/)
copy(s[i+1:], s[i:])
s[i] = x
```

### *InsertVector*
下面代码演示插入向量（封装了动态大小数组的顺序容器）的实现：
```go
a = append(a[:i], append(b, a[i:]...)...)
func Insert(s []int, k int, vs ...int) []int {
 if n := len(s) + len(vs); n <= cap(s) {
  s2 := s[:n]
  copy(s2[k+len(vs):], s[k:])
  copy(s2[k:], vs)
  return s2
 }
 s2 := make([]int, len(s) + len(vs))
 copy(s2, s[:k])
 copy(s2[k:], vs)
 copy(s2[k+len(vs):], s[k:])
 return s2
}

a = Insert(a, i, b...)
```

### *Push*
```go
a = append(a, x)
```

### *Pop*
```go
 x, a = a[len(a)-1], a[:len(a)-1]
```

### *Push Front/Unshift*
```go
a = append([]T{x}, a...)
```

### *Pop Front/Shift*
```go
x, a = a[0], a[1:]
```




## ·切片额外技巧
### *Filtering without allocating*
下面例子演示数据过滤的时候，b基于原来的a存储空间来操作，并没有重新生成新的存储空间。
```go
b := a[:0]
for _, x := range a {
 if f(x) {
  b = append(b, x)
 }
}
```
为了让截取之后没有使用的存储被GC掉，需要设置成nil：
```go
for i := len(b); i < len(a); i++ {
 a[i] = nil // or the zero value of T
}
```

### *Reversing*
反转操作演示：
```go
for i := len(a)/2-1; i >= 0; i-- {
 opp := len(a)-1-i
 a[i], a[opp] = a[opp], a[i]
}
```
还有一种方法：
```go
for left, right := 0, len(a)-1; left < right; left, right = left+1, right-1 {
 a[left], a[right] = a[right], a[left]
}
```

### *Shuffling*
洗牌算法。算法思想就是从原始数组中随机抽取一个新的数字到新数组中。
```go
for i := len(a) - 1; i > 0; i-- {
    j := rand.Intn(i + 1)
    a[i], a[j] = a[j], a[i]
}
```
go1.10之后有内置函数Shuffle[6]



### *Batching with minimal allocation*
做批处理大的切片的时候，这个技巧可以了解下：
```js
actions := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
batchSize := 3
batches := make([][]int, 0, (len(actions) + batchSize - 1) / batchSize)

for batchSize < len(actions) {
    actions, batches = actions[batchSize:], append(batches, actions[0:batchSize:batchSize])
}
batches = append(batches, actions)

// 结果：
// [[0 1 2] [3 4 5] [6 7 8] [9]]
```

### *In-place deduplicate (comparable)*
删除有序数组中的重复项：
```go
import "sort"

in := []int{3,2,1,4,3,2,1,4,1} // any item can be sorted
sort.Ints(in)
j := 0
for i := 1; i < len(in); i++ {
 if in[j] == in[i] {
  continue
 }
 j++
 // preserve the original data
 // in[i], in[j] = in[j], in[i]
 // only set what is required
 in[j] = in[i]
}
result := in[:j+1]
fmt.Println(result) // [1 2 3 4]
```


### *Sliding Window*
下面实现根据size的滑动窗口输出：
```go
func slidingWindow(size int, input []int) [][]int {
 // returns the input slice as the first element
 if len(input) <= size {
  return [][]int{input}
 }

 // allocate slice at the precise size we need
 r := make([][]int, 0, len(input)-size+1)

 for i, j := 0, size; j <= len(input); i, j = i+1, j+1 {
  r = append(r, input[i:j])
 }

 return r
}

func TestSlidingWindow(t *testing.T) {
 result := slidingWindow(2, []int{1, 2, 3, 4, 5})
 fmt.Println(result) // [[1 2] [2 3] [3 4] [4 5]]
}
```



## ·总结
1、Go切片本质上是一个结构体，保存了其长度、底层数组的容量、底层数组的指针。
2、Go切片创建方式比较多样：变量声明、切片字面量、make创建、new创建、从切片/数组截取。
3、Go切片使用len()计算长度、cap()计算容量、append()来添加元素。
4、Go切片相比数组更灵活，有很多技巧，也正因为灵活，容易发生类似内存泄露的问题，需要注意。

# Golang 常用的转义字符
1） \t : 一个制表位，实现对齐功能
2） \n : 换行符
3） \\ ： 一个\
4） \" : 一个"
5） \r : 一个回车    


# 常见的值类型和引用类型
*1)值类型*: 基本数据类型int系列, fl0at系列, bool, string、数组和结构体struct
*2)引用类型*: 指针、slice切片、map、管道chan、interface等都是引用类型

# 在 golang 中 没有前置++，--，只能后置


# 原码、反码、补码
对于*有符号*的而言:
1)二进制的最高位是符号位:0表示正数,1表示负数1 ===》[0000 0001] -1===>[1000 0001]
2)*正数*的原码，反码，补码都一样
3)*负数*的反码=它的原码符号位不变，其它位取反(0->1,1->0)
    1===>原码  [0000 0001]  反码[0000 0001]  补码[0000 0001]
    -1===>原码 [1000 0001]  反码[1111 1110]  补码[1111 1111]
4)负数的补码 = 它的反码+1
5)0的反码,补码都是0
6)在计算机运算的时候，都是以补码的方式来运算的

# 位运算
*2&3（与）*
2 的补码 0000 0010
3 的补码 0000 0011
2&3
0000 0010 =>2

*2|3=?（或）*
2 的补码 0000 0010
3 的补码 0000 0011
2|3
0000 0011 =>3

*2^3(异或)*
2 的补码 0000 0010
3 的补码 0000 0011
2^3
0000 0001 =>1

*-2^2*（负数在计算机计算时都是使用补码运算）
-2的原码1000 0010 =》反码1111 1101 =>补码1111 1110
2的补码0000 0010
-2^2
1111 1100(补码）===》原码
1111 1100 =》反码1111 1011 =》原码1000 0100 ==  -4

# go中没有 while 和 do...while

# go包的规范
## 1）包可以取别名
## 2）推荐文件夹名就是包名（package名字可以和文件夹名不一样，但不推荐）
## 3）如果要在其他地方使用包中的函数或变量（全局函数或全局变量），需要将函数或者变量名首字母大写（允许导出）才可以使用，小写无法在外面使用。
## 4）同一个包下（一个包文件夹下的所有 .go文件）不允许相同名字的全局函数或者相同的全局变量名
## 5)如果要编译一个可执行程序文件，就需要将这个包声明为 main，即package main。这个就是一个语法规范，如果写一个库，包名可以自定义


# go build -o bin\my.exe go_code/go_study/01-基础数据类型与变量/chapter11-包名与变量规范/model
·-o： 代表输出
·bin\my.exe： 代表生成名字为 my.exe 的可执行文件到当前文件夹下bin目录中
·go_code/go_study/01-基础数据类型与变量/chapter11-包名与变量规范/model ：代表要编译的文件所在目录


# Go支持自定义数据类型
基本语法： type 自定义数据类型名 数据类型 //相当取了一个别名
如：type myint int