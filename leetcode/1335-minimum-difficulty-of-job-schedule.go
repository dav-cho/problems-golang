/**
* 1335. Minimum Difficulty of a Job Schedule (hard)
**/

package main

import (
	"math"
)

// bottom-up dp space optimized (1d)
func minDifficulty1(jobDifficulty []int, d int) int {
	n := len(jobDifficulty)
	if n < d {
		return -1
	}

	dp := make([]float64, n)
	for i := range dp {
		dp[i] = math.MaxFloat64
	}
	dp[n-1] = float64(jobDifficulty[n-1])

	for i := n - 2; i >= 0; i-- {
		dp[i] = math.Max(dp[i+1], float64(jobDifficulty[i]))
	}

	for day := d - 1; day > 0; day-- {
		for i := day - 1; i < n-(d-day); i++ {
			dp[i] = math.MaxFloat64
			h := float64(0)
			for j := i; j < n-(d-day); j++ {
				h = math.Max(h, float64(jobDifficulty[j]))
				dp[i] = math.Min(dp[i], h+dp[j+1])
			}
		}
	}

	return int(dp[0])
}

// bottom-up dp (2d graph)
func minDifficulty2(jobDifficulty []int, d int) int {
	n := len(jobDifficulty)
	if n < d {
		return -1
	}

	dp := make([][]float64, n)
	for i := range dp {
		dp[i] = make([]float64, d+1)
		for j := range dp[i] {
			dp[i][j] = math.MaxFloat64
		}
	}
	dp[n-1][d] = float64(jobDifficulty[n-1])

	for i := n - 2; i >= 0; i-- {
		dp[i][d] = math.Max(dp[i+1][d], float64(jobDifficulty[i]))
	}

	for day := d - 1; day > 0; day-- {
		for i := day - 1; i < n-(d-day); i++ {
			h := float64(0)
			for j := i; j < n-(d-day); j++ {
				h = math.Max(h, float64(jobDifficulty[j]))
				dp[i][day] = math.Min(dp[i][day], h+dp[j+1][day+1])
			}
		}
	}

	return int(dp[0][1])
}
