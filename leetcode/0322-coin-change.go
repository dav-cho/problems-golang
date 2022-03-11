/**
* 322. Coin Change (medium)
**/

package main

import (
	"math"
)

// bottom-up dp
func coinChange1(coins []int, amount int) int {
	min := func(x, y int) int {
		if x < y {
			return x
		}

		return y
	}

	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0

	for _, c := range coins {
		for r := c; r <= amount; r++ {
			dp[r] = min(dp[r], 1+dp[r-c])
		}
	}

	if dp[amount] != math.MaxInt32 {
		return dp[amount]
	}

	return -1
}

// top-down dp (TLE)
func coinChange2(coins []int, amount int) int {
	memo := make(map[int]int)

	var dp func(rem int) int
	dp = func(rem int) int {
		if rem < 0 {
			return -1
		}
		if rem == 0 {
			return 0
		}
		if v, ok := memo[rem]; ok {
			return v
		}

		minCount := math.MaxInt
		for _, c := range coins {
			count := dp(rem - c)
			if count == -1 {
				continue
			}
			minCount = int(math.Min(float64(minCount), float64(count+1)))
		}

		if minCount < math.MaxInt {
			return int(minCount)
		}

		return -1
	}

	return dp(amount)
}
