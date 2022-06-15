package main

import (
	"fmt"
	"runtime"
)

func main() {
	// Go 代码在运行时检测版本
	fmt.Println("%s", runtime.Version())
}

// 你可以通过在终端输入指令 go version 来打印 Go 的版本信息。
// 也可以通过代码运行时查看 ↑
