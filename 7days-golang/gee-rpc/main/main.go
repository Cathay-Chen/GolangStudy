// Copyright 2022 Cathay.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"fmt"
	"geerpc"
	"geerpc/codec"
	"log"
	"net"
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

	// 建立 TCP 连接
	conn, _ := net.Dial("tcp", <-addr)

	defer func() {
		_ = conn.Close()
	}()

	time.Sleep(1 * time.Second)
	_ = json.NewEncoder(conn).Encode(geerpc.DefaultOption)
	cc := codec.NewGobCodec(conn)

	for i := 0; i < 5; i++ {
		h := &codec.Header{
			ServiceMethod: "Foo.Sum",
			Seq:           uint64(i),
		}

		_ = cc.Write(h, fmt.Sprintf("geerpc req %d", h.Seq))
		_ = cc.ReadHeader(h)
		var reply string
		_ = cc.ReadBody(&reply)
		log.Println("reply:", reply)
	}
}

// 运行结果：
/**
GOROOT=/usr/local/opt/go/libexec #gosetup
GOPATH=/Users/cathay/go #gosetup
/usr/local/opt/go/libexec/bin/go build -o /private/var/folders/9s/b1bl7w_x0vv9w_pbscjgcflm0000gn/T/GoLand/___1go_build_geerpc . #gosetup
/private/var/folders/9s/b1bl7w_x0vv9w_pbscjgcflm0000gn/T/GoLand/___1go_build_geerpc
start rpc server on [::]:58973
&{Foo.Sum 0 } geerpc req 0
reply: geerpc resp 0
&{Foo.Sum 1 } geerpc req 1
reply: geerpc resp 1
&{Foo.Sum 2 } geerpc req 2
reply: geerpc resp 2
&{Foo.Sum 3 } geerpc req 3
reply: geerpc resp 3
&{Foo.Sum 4 } geerpc req 4
reply: geerpc resp 4
*/
