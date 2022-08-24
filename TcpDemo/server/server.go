// Copyright 2022 Cathay.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"bufio"
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close() // 想要编辑器不显示报错，可以这样：`func() { _ = conn.Close() }()`

	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:]) // 读取数据

		if err != nil {
			fmt.Println("read from client failed, err: ", err)
			break
		}

		recvBuf := buf[:n]
		recvStr := string(recvBuf)
		fmt.Println("收到 Client 端发来的数据：", recvStr)
		conn.Write(recvBuf) // 发送数据
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:9988")

	if err != nil {
		fmt.Println("Listen() failed, err: ", err)
		return
	}

	for {
		conn, err := listen.Accept() // 监听客户端连接请求

		if err != nil {
			fmt.Println("Accept() failed, err: ", err)
			return
		}

		go process(conn)
	}
}
