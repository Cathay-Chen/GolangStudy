// 在 Go 中，变量 需要显式声明，并且在函数调用等情况下， 编译器会检查其类型的正确性。
package main

import "fmt"

func main() {
	var a = "initial"
	fmt.Println(a)

	// var 声明 1 个或者多个变量。
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Go 会自动推断已经有初始值的变量的类型。
	var d = true
	fmt.Println(d)

	// 声明后却没有给出对应的初始值时，变量将会初始化为 零值 。 例如，int 的零值是 0。
	var e int
	fmt.Println(e)

	// := 语法是声明并初始化变量的简写
	f := "short"
	fmt.Println(f)
}

// go run variables.go
// 输出：
// initial
// 1 2
// true
// 0
// short
