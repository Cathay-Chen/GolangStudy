// Copyright 2022 Cathay.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"net"
)

type User struct {
	Name string
	Addr string
	C    chan string
	conn net.Conn
}

func NewUser(conn net.Conn) *User {
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C:    make(chan string),
		conn: conn,
	}

	// 启动监听当前 User.C channel 消息的 goroutine
	go user.ListenMessage()

	return user
}

// ListenMessage 监听当前 User.C channel 的方法，一旦有消息，就直接发送给对端客户端
func (user *User) ListenMessage() {
	for {
		msg := <-user.C
		_, err := user.conn.Write([]byte(msg + "\n"))
		if err != nil {
			fmt.Println("user conn write error: ", err)
			return
		}
	}
}
