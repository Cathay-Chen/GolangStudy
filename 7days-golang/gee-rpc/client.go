// Copyright 2022 Cathay.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package geerpc

import (
	"encoding/json"
	"errors"
	"fmt"
	"geerpc/codec"
	"io"
	"log"
	"net"
	"sync"
)

// Call 代表一个执行的 RPC
type Call struct {
	Seq           uint64     // Seq
	ServiceMethod string     // ServiceMethod 格式："{service}.{method}"
	Args          any        // Args method 函数对应的参数
	Reply         any        // Reply method 函数的返回
	Error         error      // Error 如果发生错误，将被设置值
	Done          chan *Call // Done 标记完成
}

// done 标记完成
func (call *Call) done() {
	call.Done <- call
}

// Client 代表一个 RPC 客户端。可能有多个未完成的 Call。使用单个客户端，并且客户端可以由同时有多个 goroutine
type Client struct {
	cc      codec.Codec      // cc 所属 codec.Codec 子类 是消息的编解码器，和服务端类似，用来序列化将要发送出去的请求，以及反序列化接收到的响应。
	opt     *Option          // opt 请求 Option
	sending sync.Mutex       // sending 发送并发互斥锁 sending 是一个互斥锁，和服务端类似，为了保证请求的有序发送，即防止出现多个请求报文混淆。
	header  codec.Header     // header 请求的 codec.Header 是每个请求的消息头，header 只有在请求发送时才需要，而请求发送是互斥的，因此每个客户端只需要一个，声明在 Client 结构体中可以复用。
	mu      sync.Mutex       // mu 互斥锁
	seq     uint64           // seq 用于给发送的请求编号，每个请求拥有唯一编号。
	pending map[uint64]*Call // pending 存储未处理完的请求，键是编号，值是 Call 实例。

	// closing 和 shutdown 任意一个值置为 true，则表示 Client 处于不可用的状态。
	// 但有些许的差别，closing 是用户主动关闭的，即调用 Close 方法，而 shutdown 置为 true 一般是有错误发生。
	closing  bool // closing 用户主动关闭
	shutdown bool // shutdown 执行过程错误导致的停止
}

// 检查 Client 是否实现了 io.Closer 接口
var _ io.Closer = (*Client)(nil)

var ErrShutdown = errors.New("connection is shutdown")

// Close 关闭连接
func (client *Client) Close() error {
	client.mu.Lock()
	defer client.mu.Unlock()

	if client.closing {
		return ErrShutdown
	}

	client.closing = true
	return client.cc.Close()
}

// IsAvailable client 是否在工作
func (client *Client) IsAvailable() bool {
	client.mu.Lock()
	defer client.mu.Unlock()
	return !client.shutdown && !client.closing
}

// registerCall 将参数 Call 添加到 Client.pending 中，并更新 Client.seq。
func (client *Client) registerCall(call *Call) (uint64, error) {
	client.mu.Lock()
	defer client.mu.Unlock()

	if client.closing || client.shutdown {
		return 0, ErrShutdown
	}

	call.Seq = client.seq
	client.pending[call.Seq] = call
	client.seq++
	return call.Seq, nil
}

// removeCall 根据 seq，从 Client.pending 中移除对应的 Call，并返回。
func (client *Client) removeCall(seq uint64) *Call {
	client.mu.Lock()
	defer client.mu.Unlock()
	call := client.pending[seq]
	delete(client.pending, seq)
	return call
}

// terminateCalls 服务端或客户端发生错误时调用，将 shutdown 设置为 true，且将错误信息通知所有 pending 状态的 Call。
func (client *Client) terminateCalls(err error) {
	client.sending.Lock()
	defer client.sending.Unlock()
	client.mu.Lock()
	defer client.mu.Unlock()
	client.shutdown = true
	for _, call := range client.pending {
		call.Error = err
		call.done()
	}
}

// send 发送消息
func (client *Client) send(call *Call) {
	client.sending.Lock()
	defer client.sending.Unlock()

	seq, err := client.registerCall(call)
	if err != nil {
		call.Error = err
		call.done()
		return
	}

	client.header.ServiceMethod = call.ServiceMethod
	client.header.Seq = seq
	client.header.Error = ""

	if err := client.cc.Write(&client.header, call.Args); err != nil {
		call := client.removeCall(seq)

		if call != nil {
			call.Error = err
			call.done()
		}
	}
}

// receive 接受消息
// 对一个客户端端来说，接收响应、发送请求是最重要的 2 个功能。那么首先实现接收功能，接收到的响应有三种情况：
//	- Call 不存在，可能是请求没有发送完整，或者因为其他原因被取消，但是服务端仍旧处理了。
//	- Call 存在，但服务端处理出错，即 h.Error 不为空。
//	- Call 存在，服务端处理正常，那么需要从 body 中读取 Reply 的值。
func (client *Client) receive() {
	var err error
	for err == nil {
		var h codec.Header
		if err = client.cc.ReadHeader(&h); err != nil {
			break
		}
		client.header = h
		call := client.removeCall(h.Seq)
		switch {
		case call == nil: // Call 不存在，可能是请求没有发送完整，或者因为其他原因被取消，但是服务端仍旧处理了。
			err = client.cc.ReadBody(nil)
		case h.Error != "": // Call 存在，但服务端处理出错，即 h.Error 不为空。
			call.Error = fmt.Errorf(h.Error)
			err = client.cc.ReadBody(nil)
			call.done()
		default: // Call 存在，服务端处理正常，那么需要从 body 中读取 Reply 的值。
			err = client.cc.ReadBody(call.Reply)
			if err != nil {
				call.Error = errors.New("reading body " + err.Error())
			}
			call.done()
		}
	}
	// 遇到错误就停止所有请求
	client.terminateCalls(err)
}

// Go 和 Call 是客户端暴露给用户的两个 RPC 服务调用接口
// Go 是一个异步接口，返回 call 实例。
func (client *Client) Go(serviceMethod string, args, reply any, done chan *Call) *Call {
	if done == nil {
		done = make(chan *Call, 10)
	} else if cap(done) == 0 {
		log.Panic("rpc client: done channel is unbuffered.")
	}

	call := &Call{
		ServiceMethod: serviceMethod,
		Args:          args,
		Reply:         reply,
		Done:          done,
	}
	client.send(call)
	return call
}

// Call 是对 Go 的封装，阻塞 Call.Done，等待响应返回，是一个同步接口。
func (client *Client) Call(serviceMethod string, args, replay any) error {
	call := <-client.Go(serviceMethod, args, replay, make(chan *Call, 1)).Done
	log.Println("返回的 header 的内容：", call)
	return call.Error
}

// parseOptions 解析 Option， 验证参数，并赋值默认值
func parseOptions(opts ...*Option) (*Option, error) {

	if len(opts) == 0 || opts[0] == nil {
		return DefaultOption, nil
	}

	if len(opts) != 1 {
		return nil, errors.New("number of options is more than 1")
	}

	opt := opts[0]
	opt.MagicNumber = DefaultOption.MagicNumber

	if opt.CodecType == "" {
		opt.CodecType = DefaultOption.CodecType
	}

	return opt, nil
}

// NewClient 创建 Client
// 创建 Client 实例时，首先需要完成一开始的协议交换，即发送 Option 信息给服务端。
// 协商好消息的编解码方式之后，再创建一个子协程调用 receive() 接收响应。
func NewClient(conn net.Conn, opt *Option) (*Client, error) {
	f := codec.NewCodecFuncMap[opt.CodecType]

	if f == nil {
		err := fmt.Errorf("invalid codec type %s", opt.CodecType)
		log.Println("rpc client: codec error: ", err)
		_ = conn.Close()
		return nil, err
	}

	if err := json.NewEncoder(conn).Encode(opt); err != nil {
		log.Println("rpc client: options error: ", err)
		_ = conn.Close()
		return nil, err
	}

	return newClientCodec(f(conn), opt), nil
}

// newClientCodec 创建消息编码，并创建一个子携程调用 receive() 接收响应。
func newClientCodec(cc codec.Codec, opt *Option) *Client {
	client := &Client{
		seq:     1,
		cc:      cc,
		opt:     opt,
		pending: make(map[uint64]*Call),
	}
	go client.receive()
	return client
}

// Dial 实现 Dial 函数，便于用户传入服务端地址，创建 Client 实例。
// 为了简化用户调用，通过 ...*Option 将 Option 实现为可选参数。
func Dial(network, address string, opts ...*Option) (client *Client, err error) {
	opt, err := parseOptions(opts...)

	if err != nil {
		return nil, err
	}

	conn, err := net.Dial(network, address)

	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			_ = conn.Close()
		}
	}()

	return NewClient(conn, opt)
}
