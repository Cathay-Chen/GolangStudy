package gee

import (
	"net/http"
)

// HandlerFunc 路由匹配成功后执行的方法
// type HandlerFunc func(w http.ResponseWriter, r *http.Request)
type HandlerFunc func(c *Context)

// Engine 引擎
type Engine struct {
	// router map[string]HandlerFunc
	router *router
}

// ServeHTTP 实现 http.Handler interface 中的 http.Handler.ServeHTTP 方法
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	//key := req.Method + "-" + req.URL.Path
	//
	//if handler, ok := engine.router[key]; ok {
	//	handler(w, req)
	//} else {
	//	fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	//}

	c := newContext(w, req)
	engine.router.handle(c)
}

// New 初始化一个引擎
func New() *Engine {
	// return &Engine{router: make(map[string]HandlerFunc)}
	return &Engine{router: newRouter()}
}

// addRoute 添加路由
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	//key := method + "-" + pattern
	//engine.router[key] = handler
	engine.router.addRoute(method, pattern, handler)
}

// GET 添加一个 get 请求路由
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// POST 添加一个 post 请求
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// Run 启动一个 http 服务
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}
