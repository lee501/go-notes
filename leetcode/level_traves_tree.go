package main

import "fmt"

var res [][]int

func main() {
	res = [][]int{}
	head := &TreeNode{}
	levelTree(head, 0)
	fmt.Println(res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
	 3
   / \
  9  20
    /  \
   15   7
*/
func levelTree(head *TreeNode, level int) {
	if head != nil {
		//当前层初始化
		if len(res) == level {
			res = append(res, []int{})
		}
		res[level] = append(res[level], head.Val)
		levelTree(head.Left, level+1)
		levelTree(head.Right, level+1)
	}
}
