// Copyright 2022 Cathay.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"geerpc"
	"log"
	"net"
	"sync"
	"time"
)

// startServer 开启服务端服务
func startServer(addr chan string) {
	l, err := net.Listen("tcp", ":0")

	if err != nil {
		log.Fatal("network error:", err)
	}

	log.Println("start rpc server on", l.Addr())
	addr <- l.Addr().String()
	geerpc.Accept(l) // 监听
}

func main() {
	log.SetFlags(0)
	addr := make(chan string)
	go startServer(addr)
	client, _ := geerpc.Dial("tcp", <-addr)
	defer func() {
		_ = client.Close()
	}()

	time.Sleep(1 * time.Second)
	// 发送请求并接收请求
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			args := fmt.Sprintf("geerpc req %d", i)
			var reply string
			if err := client.Call("Foo.Sum", args, &reply); err != nil {
				log.Fatal("call Foo.Sum err: ", err)
			}
			log.Println("返回的 rep 的内容：", reply)
		}(i)
	}

	wg.Wait()
}

// 运行结果：
/**
GOROOT=/usr/local/opt/go/libexec #gosetup
GOPATH=/Users/cathay/go #gosetup
/usr/local/opt/go/libexec/bin/go build -o /private/var/folders/9s/b1bl7w_x0vv9w_pbscjgcflm0000gn/T/GoLand/___go_build_geerpc . #gosetup
/private/var/folders/9s/b1bl7w_x0vv9w_pbscjgcflm0000gn/T/GoLand/___go_build_geerpc
start rpc server on [::]:55971
请求的 req 头和内容： &{Foo.Sum 0 } geerpc req 0
返回的 rep 头和内容： &{Foo.Sum 0 } geerpc rep 0
请求的 req 头和内容： &{Foo.Sum 1 } geerpc req 1
返回的 rep 头和内容： &{Foo.Sum 1 } geerpc rep 1
请求的 req 头和内容： &{Foo.Sum 2 } geerpc req 2
返回的 rep 头和内容： &{Foo.Sum 2 } geerpc rep 2
请求的 req 头和内容： &{Foo.Sum 3 } geerpc req 3
返回的 rep 头和内容： &{Foo.Sum 3 } geerpc rep 3
请求的 req 头和内容： &{Foo.Sum 4 } geerpc req 4
返回的 rep 头和内容： &{Foo.Sum 4 } geerpc rep 4

Process finished with the exit code 0
*/
