package gee

import "strings"

type node struct {
	children []*node
	pattern  string // 待匹配路由，例如 /p/:lang
	part     string // 路由中的一部分，例如 :lang
	isWild   bool   // 是否精准匹配，：，*就是true
}

// 第一个匹配的节点，用于插入
func (n *node) matchChild(part string) *node {
	for _, v := range n.children {
		if v.part == part || v.isWild {
			return v
		}
	}

	return nil
}

// 所有匹配的节点，用于查找
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, v := range n.children {
		if part == v.part || v.isWild {
			nodes = append(nodes, v)
		}
	}

	return nodes
}

// 插入
func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	child := n.matchChild(parts[height])
	if child == nil {
		child = &node{
			children: nil,
			part:     parts[height],
			isWild:   parts[height][0] == ':' || parts[height][0] == '*',
		}

		n.children = append(n.children, child)
	}

	child.insert(pattern, parts, height+1)
	return
}

// 查找
func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}

		return n
	}

	part := parts[height]
	children := n.matchChildren(part)

	for _, c := range children {
		res := c.search(parts, height+1)
		if res != nil {
			return res
		}
	}

	return nil
}
