// 包声明
package main

// 引入包声明
import (
	"fmt"
	"os"
)

// 函数声明
func printInConsole(s string) {
	fmt.Println(s)
}

// 全局变量声明
var str string = "Hello world!"

func main() {
	// 1.1.0 环境配置
	traing1110()
	// 1.1.1 基本结构
	printInConsole(str)
}

func traing1110() {
	fmt.Println(os.Args)
	var s string = "Hello, World!"
	var b byte = s[0]
	fmt.Println(b)
	// fmt.Println("Hello, World!")
}
