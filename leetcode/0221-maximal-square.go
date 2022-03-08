/**
* 221. Maximal Square (medium)
**/

package main

import (
	"math"
)

// optimized dp
func maximalSquare1(matrix [][]byte) int {
	var (
		prev float64 = 0
		res  float64 = 0
	)
	m, n := len(matrix), len(matrix[0])
	dp := make([]float64, n+1)

	for r := 1; r <= m; r++ {
		for c := 1; c <= n; c++ {
			temp := dp[c]

			if matrix[r-1][c-1] == '1' {
				dp[c] = 1 + math.Min(math.Min(dp[c], dp[c-1]), prev)
				res = math.Max(res, dp[c])
			} else {
				dp[c] = 0
			}

			prev = temp
		}
	}

	return int(res * res)
}

// dp
func maximalSquare2(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	dp := make([][]float64, m+1)
	for i := range dp {
		dp[i] = make([]float64, n+1)
	}

	var res float64

	for r := 1; r <= m; r++ {
		for c := 1; c <= n; c++ {
			if matrix[r-1][c-1] == '1' {
				above := dp[r-1][c]
				left := dp[r][c-1]
				aboveLeft := dp[r-1][c-1]

				dp[r][c] = 1 + math.Min(math.Min(above, left), aboveLeft)
				res = math.Max(res, dp[r][c])
			}
		}
	}

	return int(res * res)
}
