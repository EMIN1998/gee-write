package gee

import (
	"fmt"
	"net/http"
)

type Engine struct {
	router map[string]HandleFunc
}

type HandleFunc func(w http.ResponseWriter, r *http.Request)

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	k := fmt.Sprintf("%s-%s",req.Method, req.URL.Path)
	if v, ok := e.router[k]; !ok {
		fmt.Fprint(w, "404")
	} else {
		v(w, req)
	}
}

func New() *Engine{
	return &Engine{router: make(map[string]HandleFunc)}
}

const (
	GET = "GET"
	POST = "POST"
)

func (e *Engine)addRouter(method, path string, handler HandleFunc)  {
	key := fmt.Sprintf("%s-%s", method, path)
	e.router[key] = handler
}

func (e *Engine)GET(path string, handleFunc HandleFunc)  {
	e.addRouter(GET, path, handleFunc)
}

func (e *Engine)POST(path string, handleFunc HandleFunc)  {
	e.addRouter(POST, path, handleFunc)
}

func (e *Engine) RUN(addr string) error {
	return http.ListenAndServe(addr, e)
}