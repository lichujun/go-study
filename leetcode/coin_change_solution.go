package leetcode

import "math"

func CoinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	} else if amount < 0 {
		return -1
	}
	var dp = make([]int, amount+1)
	length := len(coins)
	for i := 1; i <= amount; i++ {
		for j := 0; j < length; j++ {
			if j == 0 {
				dp[i] = math.MaxInt32
			}
			if i < coins[j] {
				continue
			}
			dp[i] = min(dp[i], dp[i-coins[j]]+1)
		}
	}
	if dp[amount] > amount {
		return -1
	} else {
		return dp[amount]
	}
}

func min(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
