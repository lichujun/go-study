package algorithm

import (
	"fmt"
	"math"
	"testing"
)

func TestCountPrimes(t *testing.T) {
	fmt.Println(countPrimes(10))
}

func countPrimes(n int) int {
	dp := make([]bool, n)
	maxSqrt := int(math.Ceil(math.Sqrt(float64(n))))
	for i := 2; i <= maxSqrt; i++ {
		if !dp[i] {
			for j := i * i; j < n; j += i {
				dp[j] = true
			}
		}
	}
	count := 0
	for _, d := range dp[2:] {
		if !d {
			count++
		}
	}
	return count
}
