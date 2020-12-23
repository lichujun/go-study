package algorithm

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestGetRandom(t *testing.T) {
	rand.Seed(time.Now().Unix())
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
	c := 0
	/*for i := 0; i < 5_000_000; i++ {
		if GetRandom(listNode) == 5 {
			c++
		}
	}
	fmt.Println(c)*/
	for i := 0; i < 5_000_000; i++ {
		for _, m := range GetKRandom(listNode, 3) {
			if m == 5 {
				c++
			}
		}
	}
	fmt.Println(c)
}

func GetRandom(node *ListNode) int {
	if node == nil {
		return 0
	}
	i, res := 0, 0
	for node != nil {
		i++
		if rand.Intn(i) == 0 {
			res = node.Val
		}
		node = node.Next
	}
	return res
}

func GetKRandom(node *ListNode, k int) []int {
	res := make([]int, k)
	i := 0
	for ; node != nil && i < k; i++ {
		res[i] = node.Val
		node = node.Next
	}
	for node != nil {
		i++
		j := rand.Intn(i)
		if j < k {
			res[j] = node.Val
		}
		node = node.Next
	}
	return res
}
