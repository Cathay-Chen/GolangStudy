// Copyright 2022 Cathay.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

func main() {
	server := NewServer("127.0.0.1", 8888)
	server.Start()
}
