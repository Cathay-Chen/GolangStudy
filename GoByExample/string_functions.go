// 标准库的 strings 包提供了很多有用的字符串相关的函数。
// 这儿有一些用来让你对 strings 包有一个初步了解的例子。
package main

import (
	"fmt"
	ss "strings"
)

// 我们给 fmt.Println 一个较短的别名， 因为我们随后会大量的使用它。
var p = fmt.Println

func main() {
	// 这是一些 strings 中有用的函数例子。
	// 由于它们都是包的函数，而不是字符串对象自身的方法， 这意味着我们需要在调用函数时，将字符串作为第一个参数进行传递。
	// 你可以在 strings 包文档中找到更多的函数。
	p("Contains:  ", ss.Contains("test", "es"))
	p("Count:     ", ss.Count("test", "t"))
	p("HasPrefix: ", ss.HasPrefix("test", "te"))
	p("HasSuffix: ", ss.HasSuffix("test", "st"))
	p("Index:     ", ss.Index("test", "e"))
	p("Join:      ", ss.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", ss.Repeat("a", 5))
	p("Replace:   ", ss.Replace("foo", "o", "0", -1))
	p("Replace:   ", ss.Replace("foo", "o", "0", 1))
	p("Split:     ", ss.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", ss.ToLower("TEST"))
	p("ToUpper:   ", ss.ToUpper("test"))
	p()

	// 虽然不是 strings 的函数，但仍然值得一提的是，
	// 获取字符串长度（以字节为单位）以及通过索引获取一个字节的机制。
	p("Len: ", len("hello"))
	p("Char:", "hello"[1])
}

// $ go run string_functions.go
// 输出：
// Contains:   true
// Count:      2
// HasPrefix:  true
// HasSuffix:  true
// Index:      1
// Join:       a-b
// Repeat:     aaaaa
// Replace:    f00
// Replace:    f0o
// Split:      [a b c d e]
// ToLower:    test
// ToUpper:    TEST
//
// Len:  5
// Char: 101
