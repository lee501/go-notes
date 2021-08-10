package queue_level_tree

import "fmt"

type Tree struct {
	Val   int
	Left  *Tree
	Right *Tree
}

//声明队列
type Queue struct {
	Val    []*Tree
	Length int
}

func (q *Queue) Push(tree *Tree) {
	q.Val = append(q.Val, tree)
}

func (q *Queue) Pop() (t *Tree) {
	l := q.Len()
	if l == 0 {
		panic("queue is blank")
	}
	t = q.Val[0]
	if l == 1 {
		q.Val = []*Tree{}
	} else {
		q.Val = q.Val[1:]
	}
	return
}

func (q *Queue) Len() int {
	q.Length = len(q.Val)
	return q.Length
}

//层序遍历树
func levelOrder(t *Tree) {
	queue := &Queue{}
	queue.Push(t)
	for queue.Len() > 0 {
		node := queue.Pop()
		if node == nil {
			panic("node is nil")
		}
		if node == t {
			fmt.Println(node.Val)
		}
		if node.Left != nil {
			fmt.Println(node.Left.Val)
			queue.Push(node.Left)
		}
		if node.Right != nil {
			fmt.Println(node.Right.Val)
			queue.Push(node.Right)
		}
	}
}
