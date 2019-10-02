package main

/*
	连表相加算法
	input (2 -> 4 -> 3) + (5 -> 6 -> 4)
	output 7 -> 0 -> 8
*/

//定义连表
type ListNode struct {
	Val int
	Next *ListNode
}

func AddTwoNumeber(left, right *ListNode) *ListNode {
	res := &ListNode{}
	cur := 0
	for left !=nil || right != nil {
		sum := cur
		if left != nil {
			sum += left.Val
			left = left.Next
		}
		if right != nil {
			sum += right.Val
			right = right.Next
		}
		//进位
		cur = sum / 10
		res.Next = &ListNode{Val: sum % 10}
		res = res.Next
	}
	return res.Next
}
