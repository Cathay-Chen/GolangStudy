package gee

import (
	"net/http"
	"strings"
)

// HandlerFunc 路由匹配成功后执行的方法
// type HandlerFunc func(w http.ResponseWriter, r *http.Request)
type HandlerFunc func(c *Context)

// Engine implement the interface of ServeHTTP
type (
	// RouterGroup 组路由
	RouterGroup struct {
		prefix      string
		engine      *Engine
		middlewares []HandlerFunc // support middleware
		parent      *RouterGroup
	}

	// Engine 引擎
	Engine struct {
		// router map[string]HandlerFunc
		router *router
		*RouterGroup
		groups []*RouterGroup
	}
)

// New 初始化一个引擎
func New() *Engine {
	// return &Engine{router: make(map[string]HandlerFunc)}
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

// Group 创建一个分组
func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	// 需要满足 group 后可以继续 group，所以重新 new
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		engine: group.engine,
		parent: group,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

// Use 添加中间件到群结构体
func (group *RouterGroup) Use(middlewares ...HandlerFunc) {
	group.middlewares = append(group.middlewares, middlewares...)
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
	var middlewares []HandlerFunc
	for _, group := range engine.groups {
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}

	c := newContext(w, req)
	c.handlers = middlewares
	engine.router.handle(c)
}
