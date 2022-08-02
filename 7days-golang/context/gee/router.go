// Copyright 2022 Cathay.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gee

import (
	"log"
	"net/http"
)

type router struct {
	handlers map[string]HandlerFunc
}

// getHandlesKey
// @Description: 获取 router.handlers 的 key
// @receiver r
// @param method
// @param pattern
// @return string
func (r router) getHandlesKey(method string, pattern string) string {
	return method + "-" + pattern
}

// addRoute
// @Description: 添加路由
// @receiver r
// @param method
// @param pattern
// @param handler
func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	key := r.getHandlesKey(method, pattern)
	r.handlers[key] = handler
}

// handle
// @Description: 根据请求的 Context 获取请求的处理方法并执行
// @receiver r
// @param c
func (r *router) handle(c *Context) {
	key := r.getHandlesKey(c.Method, c.Path)

	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}

// newRouter
// @Description: 初始化 router
// @return *router
func newRouter() *router {
	return &router{handlers: make(map[string]HandlerFunc)}
}
