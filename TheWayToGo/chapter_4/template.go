package main

import (
	"fmt"
)

const c = "C"

var v int = 5

type T struct {
}

func init() { // initialization of package
}

func main() {
	var a int
	a = 1
	// 只能在定义正确的情况下转换成功，例如从一个取值范围较小的类型转换到一个取值范围较大的类型（例如将 int16 转换为 int32）。
	// 当从一个取值范围较大的转换到取值范围较小的类型时（例如将 int32 转换为 int16 或将 float32 转换为 int），会发生精度丢失（截断）的情况。
	// 当编译器捕捉到非法的类型转换时会引发编译时错误，否则将引发运行时错误。
	// 具有相同底层类型的变量之间可以相互转换
	var s string = string(a) // 输出 空字符串
	fmt.Printf("int to string, output: %s\n", s)

	Func1()
	// ..
	fmt.Println(a)
}

func (t T) Method1() {
	// ...
}

func Func1() { //
	fmt.Println(c, v)
	// ...
}

// 所有的结构将在这一章或接下来的章节中进一步地解释说明，但总体思路如下：
// 在完成包的 import 之后，开始对常量、变量和类型的定义或声明。
// 如果存在 init 函数的话，则对该函数进行定义（这是一个特殊的函数，每个含有该函数的包都会首先执行这个函数）。
// 如果当前包是 main 包，则定义 main 函数。
// 然后定义其余的函数，首先是类型的方法，接着是按照 main 函数中先后调用的顺序来定义相关函数，如果有很多函数，则可以按照字母顺序来进行排序。

// Go 程序的执行（程序启动）顺序如下：
//
// 1. 按顺序导入所有被 main 包引用的其它包，然后在每个包中执行如下流程：
// 2. 如果该包又导入了其它的包，则从第一步开始递归执行，但是每个包只会被导入一次。
// 3. 然后以相反的顺序在每个包中初始化常量和变量，如果该包含有 init 函数的话，则调用该函数。
// 4. 在完成这一切之后，main 也执行同样的过程，最后调用 main 函数开始执行程序。
