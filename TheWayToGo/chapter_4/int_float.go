package main

import "fmt"

// Go 也有基于架构的类型，例如：int、uint 和 uintptr。
//
// 这些类型的长度都是根据运行程序所在的操作系统类型所决定的：
//
// int 和 uint 在 32 位操作系统上，它们均使用 32 位（4 个字节），在 64 位操作系统上，它们均使用 64 位（8 个字节）。
// uintptr 的长度被设定为足够存放一个指针即可。
// Go 语言中没有 float 类型。

// 整数：
//
// int8（-128 -> 127）
// int16（-32768 -> 32767）
// int32（-2,147,483,648 -> 2,147,483,647）
// int64（-9,223,372,036,854,775,808 -> 9,223,372,036,854,775,807）
// 无符号整数：
//
// uint8（0 -> 255）
// uint16（0 -> 65,535）
// uint32（0 -> 4,294,967,295）
// uint64（0 -> 18,446,744,073,709,551,615）
// 浮点型（IEEE-754 标准）：
//
// float32（+- 1e-45 -> +- 3.4 * 1e38）
// float64（+- 5 * 1e-324 -> 107 * 1e308）
// float32 精确到小数点后 7 位，float64 精确到小数点后 15 位。
// 你应该尽可能地使用 float64，因为 math 包中所有有关数学运算的函数都会要求接收这个类型。
// int 型是计算最快的一种类型。
//
//整型的零值为 0，浮点型的零值为 0.0。

func main() {
	// Go 中不允许不同类型之间的混合使用，但是对于常量的类型限制非常少，因此允许常量之间的混合使用
	// var a int
	var b int32
	// a = 15
	// b = a + a	 // 编译错误 cannot use a + a (type int) as type int32 in assignment
	b = b + 5 // 因为 5 是常量，所以可以通过编译

	// int16 也不能够被隐式转换为 int32。
	var n int16 = 34
	var m int32
	// compiler error: cannot use n (type int16) as type int32 in assignment
	// m = n
	m = int32(n)

	fmt.Printf("32 bit int is: %d\n", m)
	fmt.Printf("16 bit int is: %d\n", n)
	// 输出
	// 32 bit int is: 34
	// 16 bit int is: 34
}
