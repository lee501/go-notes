//使用结构体实现链表
package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

//遍历
func showNode(node *Node) {
	for node != nil {
		fmt.Println(*node)
		node = node.next
	}
}

/*
func main() {
	head := new(Node)
	head.data = 1
	node1 := new(Node)
	node1.data = 2
	node2 := new(Node)
	node2.data = 3

	head.next = node1
	node1.next = node2

	showNode(head)

//	链表的头节点插入
	head1 := &Node{data: 0}
	tail := head1
	for i := 1; i <= 10; i++ {
		var node = &Node{data: i}
		node.next = tail
		tail = node
	}
	showNode(tail)
}
*/
