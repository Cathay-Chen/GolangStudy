// Copyright 2022 Cathay.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"net"
)

type Server struct {
	Ip   string
	Port int
}

// NewServer 创建一个 Server 的接口
func NewServer(ip string, port int) *Server {
	return &Server{
		Ip:   ip,
		Port: port,
	}
}

// Start 启动服务接口
func (server *Server) Start() {
	// socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", server.Ip, server.Port))
	if err != nil {
		fmt.Println("net.Listen error: ", err)
		return
	}
	// close listen socket
	defer func() { _ = listener.Close() }()

	for {
		// accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept error: ", err)
			return
		}

		// do handler
		go server.Handler(conn)
	}
}

// Handler 处理链接业务
func (server *Server) Handler(conn net.Conn) {
	// ...当前链接的业务
	fmt.Println("链接建立成功")
}
