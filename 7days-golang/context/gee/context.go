// Copyright 2022 Cathay.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// H 给 map[string]interface{} 起了一个别名 H，构建 JSON 数据时，显得更简洁。
type H map[string]interface{}

// Context 目前只包含了 http.ResponseWriter 和 *http.Request，另外提供了对 Method 和 Path 这两个常用属性的直接访问。
// 提供了访问 Query 和 PostForm 参数的方法。
// 提供了快速构造 String Data JSON HTML 响应的方法。
type Context struct {
	Writer     http.ResponseWriter
	Req        *http.Request
	Path       string
	Method     string
	StatusCode int
}

// PostForm
// @Description: FormValue 返回查询的命名组件的第一个值。
//  POST 和 PUT 正文参数优先于 URL 查询字符串值。
//  如有必要，FormValue 会调用 ParseMultipartForm 和 ParseForm 并忽略这些函数返回的任何错误。
//  如果 key 不存在，FormValue 返回空字符串。
// @receiver Context c
// @param key
// @return string
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

// Query
// @Description: 获取 GET 请求参数
// @receiver Context c
// @param key
// @return string
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

// SetStatusCode
// @Description: 设置返回的 http code
// @receiver Context c
// @param code
func (c *Context) SetStatusCode(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

// SetHeader
// @Description: 设置返回的 http headers。如果 key 存在则替换值。
// 键不区分大小写，它会被 textproto.CanonicalMIMEHeaderKey 格式化。
// 要使用非规范键，请直接分配给映射。
// @receiver c
// @param key
// @param value
func (c *Context) SetHeader(key string, value string) {
	// 在 WriteHeader() 后调用 Header().Set 是不会生效的
	c.Writer.Header().Set(key, value)
}

// String
// @Description: 根据格式说明符格式化并返回结果字符串。
// @receiver c
// @param code
// @param format
// @param values
func (c *Context) String(code int, format string, values ...any) {
	// 在 WriteHeader() 后调用 Header().Set 是不会生效的，所以需要先 SetHeader 后 SetStatusCode
	c.SetHeader("Content-Type", "text/plain")
	c.SetStatusCode(code)

	if _, err := c.Writer.Write([]byte(fmt.Sprintf(format, values...))); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

// JSON
// @Description: 返回 json
// @receiver c
// @param code
// @param obj
func (c *Context) JSON(code int, obj any) {
	c.SetHeader("Content-Type", "application/json")
	c.SetStatusCode(code)
	encoder := json.NewEncoder(c.Writer)

	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

// Data
// @Description: 将数据作为 HTTP 回复的一部分写入连接。
// @receiver c
// @param code
// @param data
func (c *Context) Data(code int, data []byte) {
	c.SetStatusCode(code)

	if _, err := c.Writer.Write(data); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

// HTML
// @Description: 返回 html
// @receiver c
// @param code
// @param html
func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.SetStatusCode(code)

	if _, err := c.Writer.Write([]byte(html)); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

// newContext
// @Description: 初始化 Context
// @param w
// @param req
// @return *Context
func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}
