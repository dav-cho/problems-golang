/**
* 198. House Robber (medium)
**/

package main

import "math"

// dp optimized
func rob(nums []int) int {
	prev, curr := 0, 0

	for _, v := range nums {
		prev, curr = curr, int(math.Max(float64(curr), float64(prev+v)))
	}

	return curr
}

// dp
func rob1(nums []int) int {
	n := len(nums)
	memo := map[int]int{n - 1: nums[n-1], n: 0}

	var a, b float64
	for i := n - 2; i >= 0; i-- {
		a, b = float64(memo[i+1]), float64(memo[i+2]+nums[i])
		memo[i] = int(math.Max(a, b))
	}

	return memo[0]
}

// recursive with memoization
func rob2(nums []int) int {
	memo := make(map[int]int)

	var helper func(i int) int
	helper = func(i int) int {
		if i >= len(nums) {
			return 0
		}
		if v, ok := memo[i]; ok {
			return v
		}

		a := float64(helper(i + 1))
		b := float64(helper(i+2) + nums[i])

		memo[i] = int(math.Max(a, b))

		return memo[i]
	}

	return helper(0)
}
