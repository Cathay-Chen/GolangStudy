// Go 提供了内建的正则表达式支持。
// 这儿有一些在 Go 中与 regexp 相关的常见用法示例。
package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	pattern := "p([a-z]+)ch"

	// 测试一个字符串是否符合一个表达式。
	match, _ := regexp.MatchString(pattern, "peach")
	fmt.Println(match) // true

	// 上面我们是直接使用字符串，
	// 但是对于一些其他的正则任务，
	// 你需要通过 Compile 得到一个优化过的 Regexp 结构体。
	r, _ := regexp.Compile(pattern)

	// 该结构体有很多方法。这是一个类似于我们前面看到的匹配测试。
	fmt.Println(r.MatchString("peach")) // true

	// 查找匹配的字符串。
	fmt.Println(r.FindString("peach punch")) // peach

	// 这个也是查找首次匹配的字符串， 但是它的返回值是，匹配开始和结束位置的索引，而不是匹配的内容。
	fmt.Println("idx:", r.FindStringIndex("peach punch")) // index: [0 5]

	// Submatch 返回完全匹配和局部匹配的字符串。
	// 例如，这里会返回匹配 p([a-z]+)ch 和 ([a-z]+) 的信息。
	fmt.Println(r.FindStringSubmatch("peach punch")) // [peach ea]

	// 类似的，这个会返回完全匹配和局部匹配位置的索引。
	fmt.Println(r.FindStringSubmatchIndex("peach punch")) // [0 5 1 3]

	// 带 All 的这些函数返回全部的匹配项， 而不仅仅是首次匹配项。例如查找匹配表达式全部的项。
	fmt.Println(r.FindAllString("peach punch pinch", -1)) // [peach punch pinch]

	// All 同样可以对应到上面的所有函数。
	fmt.Println("all:", r.FindAllStringSubmatchIndex("peach punch pinch", -1)) // all: [[0 5 1 3] [6 11 7 9] [12 17 13 15]]

	// 这些函数接收一个非负整数作为第二个参数，来限制匹配次数。
	fmt.Println(r.FindAllString("peach punch pinch", 2)) // [peach punch]

	// 上面的例子中，我们使用了字符串作为参数， 并使用了 MatchString 之类的方法。
	// 我们也可以将 String 从函数名中去掉，并提供 []byte 的参数。
	fmt.Println(r.Match([]byte("peach"))) // true

	// 创建正则表达式的全局变量时，可以使用 Compile 的变体 MustCompile 。
	// MustCompile 用 panic 代替返回一个错误 ，这样使用全局变量更加安全。
	r = regexp.MustCompile(pattern)
	fmt.Println("regexp:", r) // regexp: p([a-z]+)ch

	// regexp 包也可以用来将子字符串替换为为其它值。
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>")) // a <fruit>

	// Func 变体允许您使用给定的函数来转换匹配的文本。
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out)) // a PEACH
}
