/**
* 198. House Robber (medium)
**/

package main

import "math"

func rob(nums []int) int {
	prev, curr := 0, 0

	for _, v := range nums {
		prev, curr = curr, int(math.Max(float64(curr), float64(prev+v)))
	}

	return curr
}

func rob(nums []int) int {
	n := len(nums)
	memo := map[int]int{n - 1: nums[n-1], n: 0}

	var a, b float64
	for i := n - 2; i >= 0; i-- {
		a, b = float64(memo[i+1]), float64(memo[i+2]+nums[i])
		memo[i] = int(math.Max(a, b))
	}

	return memo[0]
}
