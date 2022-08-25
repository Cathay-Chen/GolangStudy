// Copyright 2022 Cathay.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package geerpc

import (
	"encoding/json"
	"fmt"
	"geerpc/codec"
	"io"
	"log"
	"net"
	"reflect"
	"sync"
)

const MagicNumber = 0x3bef5c

type Option struct {
	// MagicNumber 标记一个请求
	MagicNumber int
	// CodecType 消息编码类型
	CodecType codec.Type
}

var DefaultOption = &Option{
	MagicNumber: MagicNumber,
	CodecType:   codec.GobType,
}

type Server struct{}

// NewServer 初始化一个服务
func NewServer() *Server {
	return &Server{}
}

var DefaultServer = NewServer()

// request 请求类
type request struct {
	// h 头信息
	h *codec.Header

	// argv, replyv 参数信息
	// 相当于
	//		argv reflect.Value 请求的参数
	//		replyv reflect.Value 返回的参数
	argv, replyv reflect.Value
}

// readRequestHeader 读取请求头信息
func (server *Server) readRequestHeader(cc codec.Codec) (*codec.Header, error) {
	var h codec.Header
	if err := cc.ReadHeader(&h); err != nil {
		if err != io.EOF && err != io.ErrUnexpectedEOF {
			log.Println("rpc server: read header error:", err)
		}
		return nil, err
	}
	return &h, nil
}

// readRequest 读取请求中的 head 和 body 信息
func (server *Server) readRequest(cc codec.Codec) (*request, error) {
	h, err := server.readRequestHeader(cc)

	if err != nil {
		return nil, err
	}

	req := &request{h: h}

	// TODO 参数类型位置，目前只处理字符串类型
	req.argv = reflect.New(reflect.TypeOf(""))

	if err = cc.ReadBody(req.argv.Interface()); err != nil {
		log.Println("rpc server: read argv err:", err)
	}

	return req, nil
}

// sendResponse 返回数据
func (server *Server) sendResponse(cc codec.Codec, h *codec.Header, body any, sending *sync.Mutex) {
	sending.Lock()
	defer sending.Unlock()

	if err := cc.Write(h, body); err != nil {
		log.Println("rpc server: write response error:", err)
	}
}

// handleRequest 处理请求
func (server *Server) handleRequest(cc codec.Codec, req *request, sending *sync.Mutex, wg *sync.WaitGroup) {
	// todo 应调用已注册的 rpc 方法以获取正确的 replyv ，目前只需要打印 argv 并发送 hello 消息
	defer wg.Done()
	log.Println("请求的 req 头和内容：", req.h, req.argv.Elem())
	req.replyv = reflect.ValueOf(fmt.Sprintf("geerpc rep %d", req.h.Seq))
	server.sendResponse(cc, req.h, req.replyv.Interface(), sending)
}

// serveCodec 根据编码类处理请求内容
func (server *Server) serveCodec(cc codec.Codec) {
	sending := new(sync.Mutex)
	wg := new(sync.WaitGroup)

	for {
		req, err := server.readRequest(cc)
		if err != nil {
			if req == nil {
				break
			}

			req.h.Error = err.Error()
			server.sendResponse(cc, req.h, req.replyv.Interface(), sending)
			continue
		}
		wg.Add(1)
		go server.handleRequest(cc, req, sending, wg)
	}

	wg.Wait()
	_ = cc.Close()
}

// invalidRequest is a placeholder for response argv when error occurs
// var invalidRequest = struct{}{}

// ServeConn 处理请求
func (server *Server) ServeConn(conn io.ReadWriteCloser) {
	defer func() {
		_ = conn.Close()
	}()

	var opt Option

	// 判断 option 格式是否错误 ， 请求时候有 encode DefaultOption
	if err := json.NewDecoder(conn).Decode(&opt); err != nil {
		log.Println("rpc server: options error:", err)
	}

	// 不知道为什么加这个逻辑
	if opt.MagicNumber != MagicNumber {
		log.Printf("rpc server: invalid magic number %x", opt.MagicNumber)
		return
	}

	// 找到编码对应的处理 Func，这个 Func 会返回处理类的实例
	f := codec.NewCodecFuncMap[opt.CodecType]
	if f == nil {
		log.Printf("rpc server: invalid codec type %s", opt.CodecType)
	}

	// 执行处理
	server.serveCodec(f(conn))
}

// Accept 开启监听并处理请求
func (server *Server) Accept(lis net.Listener) {
	for {
		conn, err := lis.Accept()

		// 如果出错就结束监听
		if err != nil {
			log.Println("rpc server: accept error:", err)
			return
		}

		// 协成处理请求
		go server.ServeConn(conn)
	}
}

// Accept new Server 监听并处理请求
func Accept(lis net.Listener) {
	DefaultServer.Accept(lis)
}
