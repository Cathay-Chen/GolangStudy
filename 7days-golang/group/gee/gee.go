package gee

import (
	"net/http"
)

// HandlerFunc 路由匹配成功后执行的方法
// type HandlerFunc func(w http.ResponseWriter, r *http.Request)
type HandlerFunc func(c *Context)

// Engine implement the interface of ServeHTTP
type (
	// RouterGroup 组路由
	RouterGroup struct {
		prefix string
		engine *Engine
	}

	// Engine 引擎
	Engine struct {
		// router map[string]HandlerFunc
		router *router
		*RouterGroup
	}
)

// New 初始化一个引擎
func New() *Engine {
	// return &Engine{router: make(map[string]HandlerFunc)}
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	return engine
}

// Group 创建一个分组
func (group *RouterGroup) Group(prefix string) *RouterGroup {
	return &RouterGroup{
		prefix: group.prefix + prefix,
		engine: group.engine,
	}
}

// addRoute 添加路由
func (group *RouterGroup) addRoute(method string, pattern string, handler HandlerFunc) {
	//key := method + "-" + pattern
	//engine.router[key] = handler
	pattern = group.prefix + pattern
	group.engine.router.addRoute(method, pattern, handler)
}

// GET 添加一个 get 请求路由
func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.engine.addRoute("GET", pattern, handler)
}

// POST 添加一个 post 请求
func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.engine.addRoute("POST", pattern, handler)
}

// Run 启动一个 http 服务
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
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
