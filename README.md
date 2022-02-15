# **Go切片与技巧**

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
1、简单的表达式
简单切片表达式的格式[low:high]。

如下面例子，n为一个切片，当用这个表达式[1:4]表示的是左闭右开[low, high)区间截取一个新的切片（例子中结果是[2 3 4]），切片被截取之后，截取的长度是high-low。

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

不满足以上条件会发送越界panic。

截取字符串
注意截取字符串，产生的也是新的字符串。

 s := "hello taibai"
 s1 := s[6:]
 fmt.Println(s1)                 // taibai
 fmt.Println(reflect.TypeOf(s1)) // string
2、扩展表达式
简单表达式生产的新切片与原数组或切片会共享底层数组，虽然避免了copy，但是会带来一定的风险。下面这个例子当新的n1切片append添加元素的时候，覆盖了原来n的索引位置4的值，导致你的程序可能是非预期的，从而产生不良的后果。

 n := []int{1, 2, 3, 4, 5, 6}
 n1 := n[1:4]
 fmt.Println(n)       // [1 2 3 4 5 6]
 fmt.Println(n1)      // [2 3 4]
 n1 = append(n1, 100) // 把n的索引位置4的值从原来的5变成了100
 fmt.Println(n)       // [1 2 3 4 100 6]
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

 n2 := n[1:4:5]         // 长度等于3，容量等于4
 fmt.Printf("%p\n", n2) // 0xc0000ac068
 n2 = append(n2, 5)
 fmt.Printf("%p\n", n2) // 0xc0000ac068
 n2 = append(n2, 6)
 fmt.Printf("%p\n", n2) // 地址发生改变，0xc0000b8000

## ·切片技巧
AppendVector
这个是添加一个切片的操作，上面我们在切片操作中已经介绍过。

a = append(a, b...)


Copy
这边给我们展示了三种copy的写法：

b := make([]T, len(a))
copy(b, a)

// 效率一般比上面的写法慢，但是如果有更多，他们效率更好
b = append([]T(nil), a...)
b = append(a[:0:0], a...)

// 这个实现等价于make+copy。
// 但在Go工具链v1.16上实际上会比较慢。
b = append(make([]T, 0, len(a)), a...)


Cut
截掉切片[i,j）之间的元素：

a = append(a[:i], a[j:]...)


Cut（GC）
上面的Cut如果元素是指针的话，会存在内存泄露，所以我们要对删除的元素设置nil，等待GC。


copy(a[i:], a[j:])
for k, n := len(a)-j+i, len(a); k < n; k++ {
 a[k] = nil // or the zero value of T
}
a = a[:len(a)-j+i]


Delete
删除索引位置i的元素：

a = append(a[:i], a[i+1:]...)
// or
a = a[:i+copy(a[i:], a[i+1:])]


Delete（GC）
删除索引位置i的元素：

copy(a[i:], a[i+1:])
a[len(a)-1] = nil // or the zero value of T
a = a[:len(a)-1]


Delete without preserving order
删除索引位置i的元素，把最后一位放到索引位置i上，然后把最后一位元素删除。这种方式底层并没有发生复制操作。

a[i] = a[len(a)-1] 
a = a[:len(a)-1]


Delete without preserving order（GC）
上面的删除操作，元素是一个指针的类型或结构体指针字段，会存在最后一个元素不能被GC掉，造成泄露，把末尾的元素设置nil，等待GC。

a[i] = a[len(a)-1]
a[len(a)-1] = nil
a = a[:len(a)-1]


Expand
这个本质上是多个append的组合操作。

a = append(a[:i], append(make([]T, j), a[i:]...)...)



Extend
用新列表扩展原来的列表

a = append(a, make([]T, j)...)


Filter (in place)
下面代码演示原地删除Go切片元素：

n := 0
for _, x := range a {
 if keep(x) {
  a[n] = x
  n++
 }
}
a = a[:n]


Insert
a = append(a[:i], append([]T{x}, a[i:]...)...)
第二个append会产生新的切片，产生一次copy，可以用以下代码方式，可免去第二次的copy：

s = append(s, 0 /* 先添加一个0值*/)
copy(s[i+1:], s[i:])
s[i] = x


InsertVector
下面代码演示插入向量（封装了动态大小数组的顺序容器）的实现：

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


Push
a = append(a, x)


Pop
 x, a = a[len(a)-1], a[:len(a)-1]


Push Front/Unshift
a = append([]T{x}, a...)


Pop Front/Shift
x, a = a[0], a[1:]





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

# 在 golang 中 i++,i--都只能独立使用
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

# 1）包可以取别名
# 2）推荐文件夹名就是包名（package名字可以和文件夹名不一样，但不推荐）
# 3）如果要在其他地方使用包中的函数或变量（全局函数或全局变量），需要将函数或者变量名首字母大写（允许导出）才可以使用，小写无法在外面使用。
# 4）同一个包下（一个包文件夹下的所有 .go文件）不允许相同名字的全局函数或者相同的全局变量名
# 5)如果要编译一个可执行程序文件，就需要将这个包声明为 main，即package main。这个就是一个语法规范，如果写一个库，包名可以自定义


# go build -o bin\my.exe go_code/go_study/01-基础数据类型与变量/chapter11-包名与变量规范/model
·-o： 代表输出
·bin\my.exe： 代表生成名字为 my.exe 的可执行文件到当前文件夹下bin目录中
·go_code/go_study/01-基础数据类型与变量/chapter11-包名与变量规范/model ：代表要编译的文件所在目录

# 基本数据类型（变量和数组）是值拷贝的方式，以值传递的方式传入函数
# （如果传入的是基本数据类型）如果希望函数内的变量能够改变函数外的变量，可以传入变量的地址&，函数内以指针的方式操作。

# Go支持自定义数据类型
基本语法： type 自定义数据类型名 数据类型 //相当取了一个别名
如：type myint int