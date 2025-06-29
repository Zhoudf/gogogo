// //---------------------------------------基本结构 Start-------------------------------
// // 包声明
// package main

// // 引入包声明
// import (
// 	"fmt"
// 	"os"
// )

// // 函数声明
// func printInConsole(s string) {
// 	fmt.Println(s)
// }

// // 全局变量声明
// var str string = "Hello world!"

// func main() {
// 	// 1.1.0 环境配置
// 	traing1110()
// 	// 1.1.1 基本结构
// 	printInConsole(str)
// }

// func traing1110() {
// 	fmt.Println(os.Args)
// 	var s string = "Hello, World!"
// 	var b byte = s[0]
// 	fmt.Println(b)
// 	// fmt.Println("Hello, World!")
// }
// //---------------------------------------基本结构 End-------------------------------

// // ---------------------------------------初始化顺序 Start-------------------------------
// package main

// import (
// 	"fmt"

// 	_ "github.com/training/gogogo/pkg1"
// )

// // 在 Go 语言里，import 语句中使用下划线 _ 导入包，这种方式被称作空白导入（blank import）。下面详细解释其作用和使用场景。
// // 作用
// // 当使用 _ 导入包时，Go 编译器只会执行该包的 init 函数，而不会将包名引入到当前作用域，也就是说你不能在代码里直接调用该包的其他函数、变量或类型。
// // 使用场景
// // 空白导入主要用于在程序启动时执行某些初始化操作，比如注册数据库驱动、注册 HTTP 路由等。许多第三方库会提供 init 函数，在 init 函数里完成一些必要的初始化工作，使用空白导入就能触发这些初始化操作。

// // 初始化顺序如下
// // pkg2.PkgNameVar has been initialized
// // pkg2 init method invoked
// // pkg1.PkgNameVar has been initialized
// // pkg1 init method invoked
// // main.getMainVar method invoked!
// // main init method invoked
// // main method invoked!

// const mainName string = "main"

// var mainVar string = getMainVar()

// func init() {
// 	fmt.Println("main init method invoked")
// }

// func main() {
// 	fmt.Println("main method invoked!")
// }

// func getMainVar() string {
// 	fmt.Println("main.getMainVar method invoked!")
// 	return mainName
// }

// //---------------------------------------初始化顺序 End-------------------------------

// ---------------------------------------基本数据类型 Start-------------------------------
// 包声明
package main

// 引入包声明
import (
	"fmt"
	"math"
)

func main() {
	// 整型：int，int8，int16，int32，int64，uint，uint8，uint16，uint32，uint64，uintptr。
	// 以 u 开头的整型被称为无符号整数类型，即都是非负数。而后面的数字代表这个值在内存中占有多少二进制位。
	// 比如 uint8 将占有 8 位，其最大值是 255
	// 比如 int8，同样占有 8 位，但是最高位是符号位，所以最大值是 127，最小值是-128（即$-2^7$）。
	// 以 uint8 举例，为整型赋值时，可以直接使用十六进制、八进制、二进制以及十进制声明。
	// int、uint 以及 uintptr 类型比较特殊，它们的值尺寸依赖于具体的编译器实现。
	// 在 64 位的架构上，它们是 64 位的。
	// 而在 32 位的架构上，它们是 32 位的。
	// 十六进制
	var a uint8 = 0xF
	var b uint8 = 0xf

	// 八进制
	var c uint8 = 017
	var d uint8 = 0o17
	var e uint8 = 0o17

	// 二进制
	var f uint8 = 0b1111
	var g uint8 = 0b1111

	// 十进制
	var h uint8 = 15
	fmt.Printf("a: %d, b: %d, c: %d, d: %d, e: %d, f: %d, g: %d, h: %d\n", a, b, c, d, e, f, g, h)

	// 	浮点数：float32，float64。
	// float32 是单精度浮点数，精确到小数点后 7 位。
	// float64 是双精度浮点数，可以精确到小数点后 15 位。
	var float1 float32 = 10
	// 在 Go 语言里，:= 是短变量声明操作符，用于在函数内部声明并初始化一个或多个变量。下面详细介绍其特性和使用场景。
	// 	特性
	// 1.声明与初始化合并：:= 操作符会同时完成变量的声明和初始化，不需要提前使用 var 关键字声明变量。
	// 2.类型推断：Go 编译器会根据赋值表达式右侧的值自动推断变量的类型，因此不需要显式指定变量类型。
	// 3.作用域：使用 := 声明的变量作用域仅限于当前代码块（通常是函数内部）。
	// 4.多变量声明：:= 可以同时声明并初始化多个变量。
	// 代码示例解释
	// 在你选中的代码 float2 := 10.0 中，:= 操作符做了以下事情：
	// 1.声明一个名为 float2 的变量。
	// 2.根据赋值表达式右侧的 10.0，Go 编译器自动推断 float2 的类型为 float64（Go 语言中，未指定类型的浮点数常量默认类型是 float64）。
	// 3.将 10.0 赋值给 float2。
	// 使用 var 关键字声明并初始化 float2
	// var float2 float64 = 10.0
	float2 := 10.0

	// 使用 fmt.Println 打印变量
	fmt.Println("float1:", float1, "float2:", float2)
	// 使用 fmt.Printf 打印变量，保留两位小数
	fmt.Printf("float1: %.2f, float2: %.2f\n", float1, float2)
	// 使用 fmt.Sprintf 生成格式化字符串
	result := fmt.Sprintf("float1: %.2f, float2: %.2f", float1, float2)
	fmt.Println(result)

	// 在实际开发中，也更推荐使用 float64 类型，因为官方标准库 math 包中，所有有关数学运算的函数的入参都是 float64 类型。
	float1 = float32(float2)
	// fmt.Printf("%v", float1 == float32(float2))
	fmt.Println(float1 == float32(float2))

	// 复数：complex64，complext128
	var c1 complex64
	c1 = 1.10 + 0.1i
	c2 := 1.10 + 0.1i        //默认complext128
	c3 := complex(1.10, 0.1) // c2与c3是等价的
	fmt.Println(c1 == complex64(c2))
	// 	在 Go 语言里，fmt.Println(complex128(c1) == c2) 返回 false，这和复数类型的精度以及浮点数比较的特性有关，下面详细解释。
	// 类型和精度差异
	// 在你的代码里，c1 是 complex64 类型，complex64 由两个 float32 组成，分别代表实部和虚部；c2 是 complex128 类型，complex128 由两个 float64 组成。float32 和 float64 的精度不同，float32 是单精度浮点数，精确到小数点后 7 位左右，而 float64 是双精度浮点数，能精确到小数点后 15 位左右。
	// 浮点数存储的近似性
	// 浮点数在计算机中是以二进制形式存储的，很多十进制小数无法用二进制精确表示，只能近似存储。在赋值 c1 = 1.10 + 0.1i 和 c2 := 1.10 + 0.1i 时，由于 c1 和 c2 的精度不同，它们存储的实际值可能存在微小差异。
	fmt.Println(complex128(c1) == c2)
	// 使用自定义函数比较
	epsilon := 1e-9
	fmt.Println(complexEqual(complex128(c1), c2, epsilon))
	fmt.Println(c2 == c3)

	// byte 是 uint8 的内置别名，可以把 byte 和 uint8 视为同一种类型。
	// 在 Go 中，字符串可以直接被转换成 []byte（byte 切片）。
	var s1 string = "Hello, world!"
	var bytes1 []byte = []byte(s1)
	fmt.Println("convert \"Hello, world!\" to bytes: ", bytes1)

	// 同时[]byte 也可以直接转换成 string。
	var bytes2 []byte = []byte{72, 101, 108, 108, 111, 44, 32, 119, 111, 114, 108, 100, 33}
	var s2 string = string(bytes2)
	fmt.Println(s2)

	// rune 是 int32 的内置别名，可以把 rune 和 int32 视为同一种类型。但 rune 是特殊的整数类型。
	// 在 Go 中，一个 rune 值表示一个 Unicode 码点。一般情况下，一个 Unicode 码点可以看做一个 Unicode 字符。有些比较特殊的 Unicode 字符有多个 Unicode 码点组成。
	var r1 rune = 'a'
	var r2 rune = '世'
	fmt.Println(r1)
	fmt.Println(r2)

	// 字符串可以直接转换成 []rune（rune 切片）。
	var s string = "abc，你好，世界！"
	var runes []rune = []rune(s)
	fmt.Println(runes)

	// 一种是解释型字面表示（interpreted string literal，双引号风格）。
	// 另一种是直白字面量表示（raw string literal， 反引号风格）。反引号中间不能加tab，要像下面顶格
	var str1 string = "Hello\nworld!\n"
	var str2 string = `Hello
world!
`
	fmt.Println(str1 == str2)

	//string,byte,rune对比
	var s3 string = "Go语言"
	var bytes3 []byte = []byte(s3)
	var runes3 []rune = []rune(s3)

	fmt.Println("string length: ", len(s3))    // 1+1+3+3=8
	fmt.Println("bytes length: ", len(bytes3)) //字节切片同string
	fmt.Println("runes length: ", len(runes3)) //4个unicode

	var s4 string = "Go语言"
	var bytes4 []byte = []byte(s4)
	var runes4 []rune = []rune(s4)

	fmt.Println("string sub: ", s4[0:7])            //从中切断乱码：Go语�
	fmt.Println("bytes sub: ", string(bytes4[0:7])) //从中切断乱码：Go语�
	fmt.Println("runes sub: ", string(runes4[0:3])) //符文切掉最后一个unicode，不乱码：Go语

	//初始值
	var digit int
	var s5 string
	var bo bool
	fmt.Println(digit)
	fmt.Println(s5)
	fmt.Println(bo)

}

// complexEqual 比较两个复数是否近似相等
func complexEqual(c1, c2 complex128, epsilon float64) bool {
	realDiff := math.Abs(real(c1) - real(c2)) //获取实数部
	imagDiff := math.Abs(imag(c1) - imag(c2)) //获取虚数部
	fmt.Println(real(c2))
	fmt.Println(imag(c2))
	return realDiff < epsilon && imagDiff < epsilon
}

//---------------------------------------基本数据类型 End-------------------------------
