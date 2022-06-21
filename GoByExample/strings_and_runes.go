// Go语言中的字符串是一个只读的byte类型的切片。
// Go语言和标准库特别对待字符串 - 作为以 UTF-8 为编码的文本容器。
// 在其他语言当中， 字符串由”字符”组成。
// 在Go语言当中，字符的概念被称为 rune - 它是一个表示 Unicode 编码的整数。
// 这个Go博客 很好的介绍了这个主题。 @link https://go.dev/blog/strings
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// s 是一个 string 分配了一个 literal value 表示泰语中的单词 “hello” 。 Go 字符串是 UTF-8 编码的文本。
	const s = "สวัสดี"

	// 因为字符串等价于 []byte， 这会产生存储在其中的原始字节的长度。
	fmt.Println("Len:", len(s))

	// 对字符串进行索引会在每个索引处生成原始字节值。 这个循环生成构成s中 Unicode 的所有字节的十六进制值。
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	// 要计算字符串中有多少rune，我们可以使用utf8包。 注意RuneCountInString的运行时取决于字符串的大小。
	// 因为它必须按顺序解码每个 UTF-8 rune。
	// 一些泰语字符由多个 UTF-8 code point 表示， 所以这个计数的结果可能会令人惊讶。
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	// range 循环专门处理字符串并解码每个 rune 及其在字符串中的偏移量。
	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	// 我们可以通过显式使用 utf8.DecodeRuneInString 函数来实现相同的迭代。
	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		// 这演示了将 rune value 传递给函数。
		examineRune(runeValue)
	}
}

func examineRune(r rune) {
	// 用单引号括起来的值是 rune literals. 我们可以直接将 rune value 与 rune literal 进行比较。
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}

// 补充：
// Golang中的单引号，更类似于C语言中的char类型，其实不能算字符串，因为只能是单个的字符。
// Golang中的双引号，才是字符串，单行的，多个字符（字母数字）。
// Golang中的反引号，类似Python的三引号，可以这行的字符串，所有转义字符将被忽略...

// go run strings_and_runes.go
// 输出：
// Len: 18
// e0 b8 aa e0 b8 a7 e0 b8 b1 e0 b8 aa e0 b8 94 e0 b8 b5
// Rune count: 6
// U+0E2A 'ส' starts at 0
// U+0E27 'ว' starts at 3
// U+0E31 'ั' starts at 6
// U+0E2A 'ส' starts at 9
// U+0E14 'ด' starts at 12
// U+0E35 'ี' starts at 15
//
// Using DecodeRuneInString
// U+0E2A 'ส' starts at 0
// found so sua
// U+0E27 'ว' starts at 3
// U+0E31 'ั' starts at 6
// U+0E2A 'ส' starts at 9
// found so sua
// U+0E14 'ด' starts at 12
// U+0E35 'ี' starts at 15
