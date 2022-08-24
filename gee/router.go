package gee

import "strings"

type routers struct {
	roots   map[string]*node
	handler map[string]HandleFunc
}

func NewRouters() *routers {
	return &routers{
		roots:   make(map[string]*node),
		handler: make(map[string]HandleFunc),
	}
}

func parsePattern(pattern string) []string {
	res := strings.Split(pattern, "/")

	resp := make([]string, 0)

	for _, s := range res {
		if s != "" {
			if s[0] == '*' {
				break
			}

			resp = append(resp, s)
		}
	}

	return resp
}
