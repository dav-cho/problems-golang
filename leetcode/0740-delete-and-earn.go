/**
* 740. Delete and Earn (medium)
**/

package main

import "math"

func deleteAndEarn(nums []int) int {
	counts := make(map[int]int)
	max := -math.MaxInt

	for _, v := range nums {
		counts[v]++

		if v > max {
			max = v
		}
	}

	points := make([]int, max+1)

	for k, v := range counts {
		points[k] = k * v
	}

	p, c := 0, 0

	for _, v := range points {
		p, c = c, int(math.Max(float64(c), float64(p+v)))
	}

	return c
}
