/**
* 1. Two Sum (easy)
**/

package main

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		comp := target - nums[i]
		if v, ok := seen[comp]; ok {
			return []int{v, i}
		}

		seen[nums[i]] = i
	}

	return []int{}
}
