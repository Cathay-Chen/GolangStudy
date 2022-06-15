package main

import "fmt"

func main() {
	// 使用 type 关键字可以定义你自己的类型，你可能想要定义一个结构体(第 10 章)，但是也可以定义一个已经存在的类型的别名，如：
	type IZ int
	var a IZ = 5
	// 这里并不是真正意义上的别名，因为使用这种方法定义之后的类型可以拥有更多的特性，且在类型转换时必须显式转换。
	// 这里我们可以看到 int 是变量 a 的底层类型，这也使得它们之间存在相互转换的可能

	println(a)

	// 多个类型需要定义，可以使用因式分解关键字的方式
	type (
		IZZ int
		FZ  float64
		STR string
	)

	var i IZZ = 1
	var f FZ = 1.22
	var s STR = "string"

	fmt.Println(i, f, s)
}
