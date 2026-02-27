package main

/*
 	  4
    /   \
   2     7
  / \   / \
 1   3 6   9
*/
//输出镜像
func transMirror(root *Tree) *Tree {
	if root != nil {
		root.Left, root.Right = root.Right, root.Left
		transMirror(root.Left)
		transMirror(root.Right)
	}
	return root
}
