package algorithm

import (
	"fmt"
	"testing"
)

func TestReverseKGroup(t *testing.T) {
	listNode := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	node := reverseKGroup(listNode, 2)
	for node != nil {
		fmt.Println(node)
		node = node.Next
	}
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	start, end := head, head
	for i := 0; i < k; i++ {
		if end == nil {
			return head
		}
		end = end.Next
	}
	newHead := reverseKNode(start, end)
	start.Next = reverseKGroup(end, k)
	return newHead
}

func reverseKNode(start *ListNode, end *ListNode) *ListNode {
	var cur, pre, next *ListNode = start, nil, nil
	for cur != end {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

type ListNode struct {
	Val  int
	Next *ListNode
}
