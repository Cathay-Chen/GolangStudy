package main

// 常量使用关键字 const 定义，用于存储**不会改变的数据**。
//
// 存储在常量中的数据类型**只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型**。

const b string = "abc" // 显式定义
const d = "abc"        // 隐式定义

func main1() {
	const Ln2 = 0.693147180559945309417232121458
	const Log2E = 1 / Ln2 // this is a precise reciprocal
	const Billion = 1e9   // float constant
	const hardEight = (1 << 100) >> 97

	// 常量还可以用作枚举
	const (
		Unknown = 0
		Female  = 1
		Male    = 2
	)

	const (
		a = iota
		b = iota
		c = iota
	)

	// 第一个 iota 等于 0，每当 iota 在新的一行被使用时，它的值都会自动加 1；所以 a=0, b=1, c=2 可以简写为如下形式：

	const (
		e = iota
		f
		g
	)

}
