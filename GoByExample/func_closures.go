// Go 支持匿名函数， 并能用其构造 闭包。
// 匿名函数在你想定义一个不需要命名的内联函数时是很实用的。
package main

import (
	"fmt"
	"reflect"
)

func main() {
	nextInt := intSeq()

	// 获取变量类型
	// 方法1：
	fmt.Println("name:", reflect.TypeOf(nextInt).Name())

	// 方法2：
	fmt.Println(reflect.TypeOf(nextInt))

	// 方法3：
	fmt.Printf("%T\n", nextInt)

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	nextInt2 := intSeq()
	fmt.Println(nextInt2())
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
