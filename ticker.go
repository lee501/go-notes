package main

import (
	"fmt"
)

type Node struct {
	Path     string
	Children []*Node
}

// retainNodes 递归地保留路径在pathsToRetain中的节点及其子节点
func retainNodes(nodes []*Node, pathsToRetain []string) []*Node {
	var retained []*Node
	pathSet := make(map[string]struct{})
	for _, path := range pathsToRetain {
		pathSet[path] = struct{}{} // 使用集合去重并快速查找
	}

	for _, node := range nodes {
		if _, exists := pathSet[node.Path]; exists {
			// 如果当前节点的路径在保留列表中，保留该节点
			retained = append(retained, node)
		} else {
			// 递归地保留子节点
			retainedChildren := retainNodes(node.Children, pathsToRetain)
			if len(retainedChildren) > 0 {
				// 如果有子节点被保留，那么当前节点也应该被保留
				node.Children = retainedChildren
				retained = append(retained, node)
			}
		}
	}
	return retained
}

func tickerDemo() {
	// 创建一个示例树
	root := &Node{
		Path: "/root",
		Children: []*Node{
			{
				Path: "/root/child1",
				Children: []*Node{
					{Path: "/root/child1/grandchild1"},
					{Path: "/root/child1/grandchild2"},
				},
			},
			{
				Path: "/root/child2",
				Children: []*Node{
					{Path: "/root/child2/grandchild1"},
				},
			},
		},
	}

	// 要保留的节点路径列表
	pathsToRetain := []string{"/root/child1/grandchild2", "/root/child2"}

	// 保留指定路径的节点及其子树
	retainedRoot := retainNodes([]*Node{root}, pathsToRetain)

	// 打印保留后的树结构
	printTree(retainedRoot, "")
}

// printTree 递归打印树结构
func printTree(nodes []*Node, prefix string) {
	for _, node := range nodes {
		fmt.Println(prefix + node.Path)
		printTree(node.Children, prefix+"  ")
	}
}
