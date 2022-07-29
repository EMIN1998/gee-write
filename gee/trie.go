package gee

type node struct {
	children []*node
	pattern  string
	part     string
	isWild   bool // 是否精准匹配，：，*就是true
}

func (n *node) matchPath(part string) *node {
	for _, v := range n.children {
		if v.part == part {
			return v
		}
	}

	return nil
}
