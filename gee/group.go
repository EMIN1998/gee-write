package gee

import (
	"net/http"
	"path"
)

type RouterGroup struct {
	engine     *Engine
	Prefix     string
	middleware []HandleFunc
	parent     *RouterGroup
}

func (r *RouterGroup) Group(prefix string) *RouterGroup {
	engine := r.engine
	newGroup := &RouterGroup{
		engine:     engine,
		Prefix:     r.Prefix + prefix,
		middleware: r.middleware,
		parent:     r,
	}

	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (r *RouterGroup) addRouter(method string, comp string, handler HandleFunc) {
	pattern := r.Prefix + comp
	//logger.Infof("%4s -- %s", method, pattern)
	r.engine.addRouter(method, pattern, handler)
}

func (r *RouterGroup) GET(path string, handleFunc HandleFunc) {
	r.addRouter(GET, path, handleFunc)
}

func (r *RouterGroup) POST(path string, handleFunc HandleFunc) {
	r.addRouter(POST, path, handleFunc)
}

func (r *RouterGroup) Use(middlewares ...HandleFunc) {
	r.middleware = append(r.middleware, middlewares...)
}

func (r *RouterGroup) createStaticHandler(filePath string, fs http.FileSystem) HandleFunc {
	absolutePath := path.Join(r.Prefix, filePath)
	fileServer := http.StripPrefix(absolutePath, http.FileServer(fs))

	return func(c *Context) {
		file := c.GetParams("filepath")
		if _, err := fs.Open(file); err != nil {
			c.Status(http.StatusNotFound)
			return
		}

		fileServer.ServeHTTP(c.Writer, c.Request)
	}
}

func (r *RouterGroup) Static(realPath string, root string) {
	handler := r.createStaticHandler(realPath, http.Dir(root))
	urlPattern := path.Join(realPath, "*filepath")
	r.GET(urlPattern, handler)
}
