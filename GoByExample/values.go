// Go 拥有多种值类型，包括字符串、整型、浮点型、布尔型等。 下面是一些基础的例子。
package main

import "fmt"

func main() {
	fmt.Println("go" + "lang")

	fmt.Println("1+1 =", 1+1) // 输出的=号后面会加一个空格
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}

// go run values.go
// 输出：
// golang
// 1+1 = 2
// 7.0/3.0 = 2.3333333333333335
// false
// true
// false
