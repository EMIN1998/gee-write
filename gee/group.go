package gee

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
