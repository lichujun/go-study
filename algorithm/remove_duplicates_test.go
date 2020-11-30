package algorithm

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	arr := []int{0, 0, 1, 1, 2, 2, 3, 3}
	fmt.Println(removeDuplicates(arr))
}

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	slow := 1
	last := nums[0]
	for _, cur := range nums[1:] {
		if cur == last {
			continue
		}
		last = cur
		nums[slow] = cur
		slow++
	}
	return slow
}
