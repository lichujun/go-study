package algorithm

import (
	"fmt"
	"testing"
)

func TestReverseList(t *testing.T) {
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
	newHead := reverseListByIteration(listNode)
	for newHead != nil {
		fmt.Println(newHead)
		newHead = newHead.Next
	}
}

func reverseListByRecursion(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseListByRecursion(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

func reverseListByIteration(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur := head.Next
	pre := head
	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	head.Next = nil
	return pre
}
