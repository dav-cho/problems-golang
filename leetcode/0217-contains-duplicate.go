/**
* 217. Contains Duplicate (easy)
**/

package main

func containsDuplicate(nums []int) bool {
	seen := make(map[int]struct{})
	for _, num := range nums {
		if _, ok := seen[num]; ok {
			return true
		}

		seen[num] = struct{}{}
	}

	return false
}
