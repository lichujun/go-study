package algorithm

import (
	"fmt"
	"testing"
)

func TestLru(t *testing.T) {
	lru := NewLru(4)
	lru.Put(1, "1")
	lru.Put(2, "2")
	lru.Put(3, "3")
	lru.Put(4, "4")
	lru.Put(5, "5")

	node := lru.head
	fmt.Println(node)
	for node.next != nil {
		fmt.Println(node.next)
		node = node.next
	}

}
