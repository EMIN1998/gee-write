package gee

import (
	logger "github.com/amoghe/distillog"
	"html/template"
	"net/http"
	"reflect"
	"runtime"
	"strings"
)

type Engine struct {
	*RouterGroup
	router        *router
	groups        []*RouterGroup
	httpTemplates *template.Template
	funcMap       template.FuncMap
}

type HandleFunc func(ctx *Context)

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var middlewares []HandleFunc
	for _, group := range e.groups {
		if strings.HasPrefix(req.URL.Path, group.Prefix) {
			middlewares = append(middlewares, group.middleware...)
		}
	}

	c := NewContext(req, w)
	c.handlers = middlewares
	c.engine = e
	e.router.handle(c)
}

func New() *Engine {
	engine := &Engine{router: NewRouters()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = make([]*RouterGroup, 0)
	return engine
}

const (
	GET  = "GET"
	POST = "POST"
)

func (e *Engine) addRouter(method, path string, handler HandleFunc) {
	e.router.addRouter(method, path, handler)
}

func (e *Engine) GET(path string, handleFunc HandleFunc) {
	e.addRouter(GET, path, handleFunc)
}

func (e *Engine) POST(path string, handleFunc HandleFunc) {
	e.addRouter(POST, path, handleFunc)
}

func (e *Engine) Println() {
	for path, handleFunc := range e.router.handlers {
		logger.Infoln(path, "\b", runFuncName(handleFunc))
	}
}

// 获取正在运行的函数名
func runFuncName(fun HandleFunc) string {
	f := runtime.FuncForPC(reflect.ValueOf(fun).Pointer()).Name()
	return f
}

func (e *Engine) RUN(addr string) error {
	e.Println()
	logger.Warningf("SERVER start ！ listen: %s\n", addr)
	return http.ListenAndServe(addr, e)
}

func (e *Engine) SetFuncMap(funcMap template.FuncMap) {
	e.funcMap = funcMap
}

func (e *Engine) LoadHTMLGlob(pattern string) {
	e.httpTemplates = template.Must(template.New("").Funcs(e.funcMap).ParseGlob(pattern))
}
