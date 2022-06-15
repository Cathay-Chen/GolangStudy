// 在 Go 文件里执行 C 语言代码
package chapter_3

// include <stdlib.h>
import "C"

func Random() int {
	return int(C.random())
}

func Seed(i int) {
	C.srandom(C.uint(i))
}
