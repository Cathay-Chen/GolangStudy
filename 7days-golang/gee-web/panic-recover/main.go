// Copyright 2022 Cathay.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"gee"
	"net/http"
)

func main() {
	r := gee.Default()

	r.GET("/", func(c *gee.Context) {
		c.String(http.StatusOK, "Hello Cathay\n")
	})

	r.GET("/panic", func(c *gee.Context) {
		names := []string{"cathay"}
		c.String(http.StatusOK, names[2])
	})

	r.Run(":9998")
}
