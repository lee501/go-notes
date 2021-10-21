package main

/*
	旋转链表， 将每个节点向右移k位置
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	tail := head
	length := 1
	//计算链表长度
	for tail.Next != nil {
		tail = tail.Next
		length++
	}
	//连成环
	tail.Next = head

	for i := 0; i < length-k%length; i++ {
		tail = tail.Next
		head = head.Next
	}
	//打断环
	tail.Next = nil
	return head
}

func reverseList(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode = nil
	for cur != nil {
		pre, cur, cur.Next = cur, cur.Next, pre
	}
	return pre
}
