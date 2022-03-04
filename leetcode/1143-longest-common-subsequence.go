/**
* 1145. Longest Common Subsequence (medium)
**/

package main

import "math"

func longestCommonSubsequence(text1 string, text2 string) int {
	R, C := len(text1), len(text2)

	if R > C {
		return longestCommonSubsequence(text2, text1)
	}

	prev := make([]int, R+1)

	for c := C - 1; c >= 0; c-- {
		curr := make([]int, R+1)

		for r := R - 1; r >= 0; r-- {
			if text1[r] == text2[c] {
				curr[r] = 1 + prev[r+1]
			} else {
				curr[r] = int(math.Max(float64(curr[r+1]), float64(prev[r])))
			}
		}

		prev = curr
	}

	return prev[0]
}
