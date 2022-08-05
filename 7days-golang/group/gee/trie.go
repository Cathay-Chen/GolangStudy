// Copyright 2022 Cathay.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// trie 单词查找树，这里作为路由匹配树

package gee

import "strings"

// node 路由节点
type node struct {
	pattern  string  // pattern 待匹配路由，例如：/p/:lang，用来匹配路由对应的 HandlerFunc
	part     string  // part 路由中的一部分，例如：:lang
	children []*node // children 子节点，例如：[p, lang]
	isWild   bool    // isWild 是否模糊匹配，part 含有 `:` 或 `*` 时为 true
}

// insert 插入路由树节点，根据`/`拆分路由，从前往后，依次检查一样的放在一个 node 里，
// 后面节点依次放入 children 里，不理解可以查看[图片](https://geektutu.com/post/gee-day3/trie_router.jpg)
func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	part := parts[height]
	child := n.matchChild(part)

	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}

	child.insert(pattern, parts, height+1)
}

// search 从树节点中找到匹配的最终节点信息
func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}

		return n
	}

	part := parts[height]
	children := n.matchChildren(part)

	for _, child := range children {
		result := child.search(parts, height+1)

		if result != nil {
			return result
		}
	}

	return nil
}

func (n *node) travel(list *[]*node) {
	if n.pattern != "" {
		*list = append(*list, n)
	}
	for _, child := range n.children {
		child.travel(list)
	}
}

// matchChild 匹配子节点
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}

	return nil
}

// matchChildren 匹配 子节点们
// children 是 child 的复数
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)

	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}

	return nodes
}
