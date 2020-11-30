package algorithm

import (
	"fmt"
	"testing"
)

func TestTrap(t *testing.T) {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(doublePointerTrap(height))
}

/**
双指针
*/
func doublePointerTrap(height []int) int {
	if len(height) <= 1 {
		return 0
	}
	left, right := 1, len(height)-2
	lMax, rMax := height[0], height[len(height)-1]
	ans := 0
	for left <= right {
		if lMax > rMax {
			ans += MaxInt2(0, rMax-height[right])
			rMax = MaxInt2(rMax, height[right])
			right--
		} else {
			ans += MaxInt2(0, lMax-height[left])
			lMax = MaxInt2(lMax, height[left])
			left++
		}
	}
	return ans
}

/**
备忘录
*/
func memoTrap(height []int) int {
	ans := 0
	lMaxArr := make([]int, len(height))
	rMaxArr := make([]int, len(height))
	for i := 1; i < len(height); i++ {
		lMaxArr[i] = MaxInt2(lMaxArr[i-1], height[i-1])
	}
	for i := len(height) - 2; i >= 0; i-- {
		rMaxArr[i] = MaxInt2(rMaxArr[i+1], height[i+1])
	}
	for i, cur := range height {
		ans += MaxInt2(0, MinInt2(lMaxArr[i], rMaxArr[i])-cur)
	}
	return ans
}

/**
暴力
*/
func violenceTrap(height []int) int {
	ans := 0
	for i, cur := range height {
		lMax, rMax := 0, 0
		for j := 0; j < i; j++ {
			lMax = MaxInt2(lMax, height[j])
		}
		if lMax == 0 {
			continue
		}
		for k := i + 1; k < len(height); k++ {
			rMax = MaxInt2(rMax, height[k])
		}
		if rMax == 0 {
			continue
		}
		ans += MaxInt2(0, MinInt2(lMax, rMax)-cur)
	}
	return ans
}

func MaxInt2(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func MinInt2(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
