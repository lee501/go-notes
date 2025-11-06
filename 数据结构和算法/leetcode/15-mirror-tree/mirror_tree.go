package main

type Tree struct {
	Val   int
	Left  *Tree
	Right *Tree
}

func isMirror(tree *Tree) bool {
	if tree == nil {
		return true
	}
	return checkTree(tree.Left, tree.Right)
}

func checkTree(left, right *Tree) bool {
	if left == nil && right == nil {
		return true
	}
	if (left == nil && right != nil) || (left != nil && right == nil) {
		return false
	}
	return left.Val == right.Val && checkTree(left.Left, left.Right) && checkTree(right.Left, right.Right)
}
