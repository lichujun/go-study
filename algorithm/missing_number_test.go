package algorithm

import (
	"fmt"
	"testing"
)

func TestMissingNumber(t *testing.T) {
	fmt.Println(missingNumber2([]int{1, 0, 2, 4}))
}

func missingNumber1(nums []int) int {
	n := len(nums)
	res := n
	for i, num := range nums {
		res ^= i ^ num
	}
	return res
}

func missingNumber2(nums []int) int {
	n := len(nums)
	res := n
	for i, num := range nums {
		res += i - num
	}
	return res
}
