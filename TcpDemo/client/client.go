// Copyright 2022 Cathay.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9988")

	if err != nil {
		fmt.Println("Connect network err: ", err)
		return
	}

	defer conn.Close()
	inputReader := bufio.NewReader(os.Stdin)

	for {
		input, _ := inputReader.ReadString('\n') // 读取用户输入
		inputInfo := strings.Trim(input, "\r\n")

		fmt.Println("inputInfo: ", inputInfo)

		if strings.ToUpper(inputInfo) == "Q" { // 输入 q 退出
			return
		}

		_, err := conn.Write([]byte(inputInfo)) // 发送数据

		if err != nil {
			return
		}

		buf := [512]byte{}
		n, err := conn.Read(buf[:])

		if err != nil {
			fmt.Println("recv failed, err: ", err)
			return
		}

		fmt.Println("服务返回数据：", string(buf[:n]))
	}

}
