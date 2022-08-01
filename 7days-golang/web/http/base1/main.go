package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)
	// 第一个参数是地址，:9999表示在 9999 端口监听。
	// 而第二个参数则代表处理所有的HTTP请求的实例，nil 代表使用标准库中的实例处理。 第二个参数，则是我们基于net/http标准库实现Web框架的入口。
	log.Fatal(http.ListenAndServe(":9999", nil))
}

// 处理输出 URL 中的 Path
func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
}

// 处理输出 URL 中的 headers
func helloHandler(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}

// 接第13行
//
// 第二个参数的类型是什么呢？
// 通过查看 net/http 的源码可以发现，Handler 是一个接口，需要实现方法 ServeHTTP ，也就是说，只要传入任何实现了 ServerHTTP 接口的实例，
// 所有的 HTTP 请求，就都交给了该实例处理了。

// 源码：
// ```go
// package http
//
// type Handler interface {
//     ServeHTTP(w ResponseWriter, r *Request)
// }
//
// func ListenAndServe(address string, h Handler) error
// ```

// base2/main 实现 http.Handler 接口
