package main

var a = "G"
var e string

func main() {
	//var ii int = 1
	//var ss string = "a"
	// print(ii == ss) Go 对于值之间的比较有非常严格的限制，只有两个类型相同的值才可以进行比较，

	n()
	m()
	n()
	print("\n")
	e = "S"
	print(e)
	f1()
}

func n() {
	print(a)
}

func m() {
	a = "O"
	print(a)
}

func f1() {
	e := "B"
	print(e)
	f2()
}

func f2() {
	print(e)
}
