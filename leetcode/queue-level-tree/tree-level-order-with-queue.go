package queue_level_tree

import "fmt"

//定义queue接口
type QueueI interface {
	//method list
	Push(item interface{})
	DeQueue() interface{}
	Size() int
	IsEmpty() bool
}

type List struct {
	elements []interface{}
}

//生成实例
func New() *List {
	return &List{}
}

func (l *List) Push(item interface{}) {
	l.elements = append(l.elements, item)
}

func (l *List) DeQueue() interface{} {
	if l.IsEmpty() {
		return nil
	}
	item := l.elements[0]
	l.elements = l.elements[1:]
	return item
}

func (l *List) Size() int {
	return len(l.elements)
}

func (l *List) IsEmpty() bool {
	return len(l.elements) == 0
}

//层序遍历
type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func (node *Node) LevelTravel() {
	if node == nil {
		return
	}
	queue := New()
	queue.Push(node)
	for !queue.IsEmpty() {
		res := queue.DeQueue().(*Node)
		fmt.Println(res.Val)
		//左树
		if res.Left != nil {
			queue.Push(res.Left)
		}
		//右树
		if res.Right != nil {
			queue.Push(res.Right)
		}
	}
}
