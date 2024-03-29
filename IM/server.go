// Copyright 2022 Cathay.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"net"
	"sync"
)

type Server struct {
	Ip   string
	Port int

	// 在线用户的列表
	OnlineMap map[string]*User
	mapLock   sync.RWMutex

	// 消息广播的 channel
	Message chan string
}

// NewServer 创建一个 Server 的接口
func NewServer(ip string, port int) *Server {
	return &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
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

	// 启动监听 Message 的 goroutine
	go server.ListenMessage()

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

// ListenMessage 监听 Message 广播消息 channel 的 goroutine，
// 一旦有消息就发送给全部的在线 User
func (server *Server) ListenMessage() {
	for {
		msg := <-server.Message

		server.mapLock.Lock()
		for _, cli := range server.OnlineMap {
			cli.C <- msg

		}
		server.mapLock.Unlock()
	}
}

// Handler 处理链接业务
func (server *Server) Handler(conn net.Conn) {
	// ...当前链接的业务
	// fmt.Println("链接建立成功")

	user := NewUser(conn)

	// 用户上线，将用户加入到 onlineMap 中
	server.mapLock.Lock()
	server.OnlineMap[user.Name] = user
	server.mapLock.Unlock()

	// 广播当前用户上线消息
	server.BroadCast(user, "已上线")

	// 当前 handler 阻塞
	select {}
}

// BroadCast 广播消息
func (server *Server) BroadCast(user *User, msg string) {
	sendMsg := fmt.Sprintf("[%s]%s:%s", user.Addr, user.Name, msg)
	server.Message <- sendMsg
}
