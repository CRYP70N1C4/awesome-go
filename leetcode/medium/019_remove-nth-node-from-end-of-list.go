package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := 0
	tmp := head
	for ; tmp != nil; tmp = tmp.Next {
		length++
	}
	if length < n {
		return nil
	} else if length == n {
		return head.Next
	}

	return head
}
