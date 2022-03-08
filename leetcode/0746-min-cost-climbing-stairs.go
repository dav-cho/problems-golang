/**
* 746. Min Cost Climbing Stairs (easy)
**/

package main

import "math"

// bottom-up dp constant space
func minCostClimbingStairs1(cost []int) (curr int) {
	prev, curr := 0, 0

	for i := 1; i < len(cost); i++ {
		prev, curr = curr, int(math.Min(float64(curr+cost[i]), float64(prev+cost[i-1])))
	}

	return
}

// recursive dp w/ memoization
func minCostClimbingStairs2(cost []int) int {
	memo := make(map[int]int)
	var helper func(i int) int

	helper = func(i int) int {
		if i >= len(cost)-1 {
			return 0
		}
		if v, ok := memo[i]; ok {
			return v
		}

		a, b := float64(helper(i+1)+cost[i]), float64(helper(i+2)+cost[i+1])
		memo[i] = int(math.Min(a, b))

		return memo[i]
	}

	return helper(0)
}
