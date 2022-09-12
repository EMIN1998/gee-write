package gee

import (
	"fmt"
	"net/http"
	"strings"
)

type router struct {
	roots    map[string]*node
	handlers map[string]HandleFunc
}

func NewRouters() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandleFunc),
	}
}

func (r *router) handle(c *Context) {
	key := fmt.Sprintf("%s-%s", c.Method, c.Path)
	if handler, ok := r.handlers[key]; !ok {
		c.String(http.StatusNotFound, "404 NOT FOUND path:%s", c.Path)
	} else {
		handler(c)
	}
}

func parsePattern(pattern string) []string {
	res := strings.Split(pattern, "/")

	resp := make([]string, 0)
	for _, s := range res {
		if s != "" {
			resp = append(resp, s)
			// 文件路径，在结尾，剩余的都是文件路径
			if s[0] == '*' {
				break
			}
		}
	}

	return resp
}

func (r *router) addRouter(method string, pattern string, handler HandleFunc) {
	parts := parsePattern(pattern)
	key := method + "-" + pattern
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}

	r.roots[method].insert(pattern, parts, 0)
	// 注册handler函数
	r.handlers[key] = handler
}

func (r *router) getRouter(method, path string) (*node, map[string]string) {
	searchPath := parsePattern(path)
	paramsMap := make(map[string]string)
	if _, ok := r.roots[method]; !ok {
		return nil, nil
	}

	targetNode := r.roots[method].search(searchPath, 0)
	if targetNode != nil {
		parts := parsePattern(targetNode.pattern)
		for index, part := range parts {
			// 获取待匹配的参数具体值
			if part[0] == ':' {
				paramsMap[part[1:]] = searchPath[index]
			}

			// 获取待匹配的文件参数具体值， 路径需要二次拼接
			if part[0] == '*' && len(part) > 1 {
				paramsMap[part[1:]] = strings.Join(searchPath[index:], "/")
			}
		}

		return targetNode, paramsMap
	}

	return nil, nil
}
