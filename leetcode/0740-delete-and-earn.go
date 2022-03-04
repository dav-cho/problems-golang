/**
* 740. Delete and Earn (medium)
**/

package main

import "math"

func deleteAndEarn(nums []int) int {
	counts := make(map[int]int)
	max := int(math.Inf(-1))
	for _, num := range nums {
		counts[num]++
		if num > max {
			max = num
		}
	}

	points := make([]int, max+1)
	for k, v := range counts {
		points[k] = k * v
	}

	prev, curr := 0, 0
	for _, p := range points {
		prev, curr = curr, int(math.Max(float64(curr), float64(prev+p)))
	}

	return curr
}
