// Copyright 2022 Cathay.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package codec

import (
	"bufio"
	"encoding/gob"
	"io"
	"log"
)

type GobCodec struct {
	conn io.ReadWriteCloser
	buf  *bufio.Writer
	dec  *gob.Decoder
	enc  *gob.Encoder
}

// 检查 GobCodec 是否实现了 Codec
var _ Codec = (*GobCodec)(nil)

func NewGobCodec(conn io.ReadWriteCloser) *GobCodec {
	buf := bufio.NewWriter(conn)
	return &GobCodec{
		conn: conn,
		buf:  buf,
		dec:  gob.NewDecoder(conn),
		enc:  gob.NewEncoder(buf),
	}
}

func (c *GobCodec) ReadHeader(h *Header) error {
	return c.dec.Decode(h)
}

func (c *GobCodec) ReadBody(body any) error {
	return c.dec.Decode(body)
}

func (c *GobCodec) Close() error {
	return c.conn.Close()
}

// Write gob 编码请求的 Header 和 body 数据
//
// ----------------------------------------------------------------------
// 这里的返回值介绍：
// 函数定义时可以给返回值命名，并在函数体中直接使用这些变量，最后通过`return`关键字返回。
func (c *GobCodec) Write(h *Header, body any) (err error) {
	defer func() {
		_ = c.buf.Flush()
		if err != nil {
			_ = c.Close()
		}
	}()

	if err = c.enc.Encode(h); err != nil {
		log.Println("rpc: god error encoding body:", err)
		return
	}

	if err = c.enc.Encode(body); err != nil {
		log.Println("rpc: god error encoding body:", err)
		return
	}
	return
}
