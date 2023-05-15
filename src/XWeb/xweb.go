package XWeb

import (
	"fmt"
	"net/http"
)

// HandlerFunc 定义了 xweb 使用的请求处理程序
type HandlerFunc func(http.ResponseWriter, *http.Request)

// Engine 实现 ServerHTTP 接口
type Engine struct {
	router map[string]HandlerFunc
}

// New 是 Engine 的构造函数
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	engine.router[key] = handler
}

// GET 定义了添加 GET 请求的方法
func (engine *Engine) GET(patten string, handler HandlerFunc) {
	engine.addRoute("GET", patten, handler)
}

// POST 定义了添加 POST 请求的方法
func (engine *Engine) POST(patten string, handler HandlerFunc) {
	engine.addRoute("POST", patten, handler)
}

// Run 定义了启动 http 服务的方法
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}
