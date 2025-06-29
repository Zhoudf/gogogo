// ---------------------------------------初始化顺序 Start-------------------------------
package main

import (
	"fmt"

	_ "github.com/training/gogogo/basic1/pkg1"
)

// 在 Go 语言里，import 语句中使用下划线 _ 导入包，这种方式被称作空白导入（blank import）。下面详细解释其作用和使用场景。
// 作用
// 当使用 _ 导入包时，Go 编译器只会执行该包的 init 函数，而不会将包名引入到当前作用域，也就是说你不能在代码里直接调用该包的其他函数、变量或类型。
// 使用场景
// 空白导入主要用于在程序启动时执行某些初始化操作，比如注册数据库驱动、注册 HTTP 路由等。许多第三方库会提供 init 函数，在 init 函数里完成一些必要的初始化工作，使用空白导入就能触发这些初始化操作。

// 初始化顺序如下
// pkg2.PkgNameVar has been initialized
// pkg2 init method invoked
// pkg1.PkgNameVar has been initialized
// pkg1 init method invoked
// main.getMainVar method invoked!
// main init method invoked
// main method invoked!

const mainName string = "main"

var mainVar string = getMainVar()

func init() {
	fmt.Println("main init method invoked")
}

func main() {
	fmt.Println("main method invoked!")
}

func getMainVar() string {
	fmt.Println("main.getMainVar method invoked!")
	return mainName
}

//---------------------------------------初始化顺序 End-------------------------------
