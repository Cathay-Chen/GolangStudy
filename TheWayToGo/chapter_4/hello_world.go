package main

// 包名 pkg ， 所有的包名都应该使用小写字母。
// 属于同一个包的源文件必须全部被一起编译，一个包即是编译时的一个单元，因此根据惯例，每个目录都只包含一个包。

import (
	"fmt"
	"strconv"
) // 标准库：在 Go 的安装文件里包含了一些可以直接使用的包，即标准库。

// 导入多个包 | 当你导入多个包时，最好按照字母顺序排列包名，这样做更加清晰易读。
// --- 1
// import "fmt"
// import "os"
// --- 2
// import "fmt"; import "os"
// --- 3
// import (
//   "fmt"
//   "os"
// )
// --- 4
// import ("fmt"; "os")

// 如果名称重复，使用别名引入
// import fm "fmt" // alias3

// 如果包名不是以 . 或 / 开头，如 "fmt" 或者 "container/list"，则 Go 会在全局文件进行查找；
// 如果包名以 ./ 开头，则 Go 会在相对目录中查找；
// 如果包名以 / 开头（在 Windows 下也可以这样使用），则会在系统的绝对路径中查找。

// 导入的包必须使用，不适用会报错  imported and not used: xx ， 这正是遵循了 Go 的格言：“没有不必要的代码！“。

// ----------------------------------------------------------------------------------------------------
// 可以在使用 import 导入包之后定义或声明 0 个或多个常量（const）、变量（var）和类型（type），
// 这些对象的作用域都是全局的（在本包范围内）， 所以可以被本包中所有的函数调用（如 gotemplate.go 源文件中的 c 和 v），
// 然后声明一个或多个函数（func）。
// ----------------------------------------------------------------------------------------------------

// main 函数是每一个可执行程序所必须包含的，一般来说都是在启动后第一个执行的函数（如果有 init() 函数则会先执行该函数）。
// 如果你的 main 包的源代码没有包含 main 函数，则会引发构建错误 undefined: main.main。
// main 函数既没有参数，也没有返回类型（与 C 家族中的其它语言恰好相反）。如果你不小心为 main 函数添加了参数或者返回类型，将会引发构建错误：
// ```
//	func main must have no arguments and no return values results.
// ```
// 在程序开始执行并完成初始化后，第一个调用（程序的入口点）的函数是 main.main()（如：C 语言），该函数一旦返回就表示程序已成功执行并立即退出。
func main2() {

	// 包内大写字母开头方法和变量名，外部可以访问。 小写字母开头的只能内部访问
	fmt.Println("hello, world")
	fmt.Printf("Καλημέρα κόσμε; or こんにちは 世界\n")
	// fmt.Print("hello, world\n") 和 fmt.Println("hello, world") 可以得到相同的结果
}

// 可以在方法名后面括号 () 中写入 0 个或多个函数的参数（使用逗号 , 分隔），每个参数的名称后面必须紧跟着该参数的类型。

// 左大括号 { 必须与方法的声明放在同一行，这是编译器的强制规定，否则你在使用 gofmt 时就会出现错误提示：
// `build-error: syntax error: unexpected semicolon or newline before {`
// **（这是因为编译器会产生 func main() ; 这样的结果，很明显这错误的）**

// ** Go 语言虽然看起来不使用分号作为语句的结束，但实际上这一过程是由编译器自动完成，因此才会引发像上面这样的错误 **

// 几乎所有全局作用域的类型、常量、变量、函数和被导出的对象都应该有一个合理的注释。如果这种注释（称为文档注释）出现在函数前面，
// 例如函数 Abcd，则要以 "Abcd..." 作为开头。
// functionName is a example
func functionName(param1 int, param2 string) string {
	return strconv.Itoa(param1) + param2
}

// 符合规范的函数一般写成如下的形式：
//
// func functionName(parameter_list) (return_value_list) {
//    …
// }
// 其中：
// parameter_list 的形式为 (param1 type1, param2 type2, …)
// return_value_list 的形式为 (ret1 type1, ret2 type2, …)
// 只有当某个函数需要被外部包调用的时候才使用大写字母开头，并遵循 Pascal 命名法；
// 否则就遵循骆驼命名法，即第一个单词的首字母小写，其余单词的首字母大写。

// 程序正常退出的代码为 0 即 Program exited with code 0；
// 如果程序因为异常而被终止，则会返回非零值，如：1。这个数值可以用来测试是否成功执行一个程序。
// ```
// 	GOROOT=/usr/local/opt/go/libexec #gosetup
//	GOPATH=/Users/cathay/go #gosetup
///	usr/local/opt/go/libexec/bin/go build -o /private/var/folders/9s/b1bl7w_x0vv9w_pbscjgcflm0000gn/T/GoLand/___go_build_hello_world_go /Users/cathay/go/src/GolangStudy/TheWayToGo/chapter_4/hello_world.go #gosetup
///	private/var/folders/9s/b1bl7w_x0vv9w_pbscjgcflm0000gn/T/GoLand/___go_build_hello_world_go
//	hello, world
//
//	Process finished with the exit code 0
// ```
