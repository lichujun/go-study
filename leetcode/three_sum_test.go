package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

func soluteThreeSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var res [][]int
	length := len(nums)
	for i := 0; i < length; {
		curNum := nums[i]
		tuples := soluteTwoSum(nums, i+1, target-curNum)
		for _, tuple := range tuples {
			tuple = append(tuple, nums[i])
			res = append(res, tuple)
		}
		for i < length && nums[i] == curNum {
			i++
		}
	}
	return res
}

func soluteTwoSum(nums []int, start int, target int) [][]int {
	length := len(nums)
	lo, hi := start, length-1
	var res [][]int
	for lo < hi {
		left, right := nums[lo], nums[hi]
		sum := left + right
		if sum < target {
			for lo < hi && nums[lo] == left {
				lo++
			}
		} else if sum > target {
			for lo < hi && nums[hi] == right {
				hi--
			}
		} else {
			res = append(res, []int{nums[lo], nums[hi]})
			for lo < hi && nums[lo] == left {
				lo++
			}
			for lo < hi && nums[hi] == right {
				hi--
			}
		}
	}
	return res
}

func TestTwoSum(t *testing.T) {
	nums := []int{-1, 0, 1, 2}
	fmt.Println(soluteTwoSum(nums, 0, 1))
}

func TestThreeSum(t *testing.T) {
	nums := []int{-4, -1, -1, 0, 1, 2}
	fmt.Println(soluteThreeSum(nums, 0))
}
