/**
* 349. Intersection of Two Arrays (easy)
**/

package main

import "sort"

// binary search
func intersection(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		return intersection(nums2, nums1)
	}

	sort.Ints(nums2)
	seen := make(map[int]struct{})

	for _, num := range nums1 {
		if search(nums2, num) {
			seen[num] = struct{}{}
		}
	}

	ret := make([]int, 0, len(seen))
	for num := range seen {
		ret = append(ret, num)
	}

	return ret
}

func search(nums []int, target int) bool {
	l, r := 0, len(nums)-1

	for l <= r {
		m := (l + r) >> 1

		switch true {
		case nums[m] == target:
			return true
		case nums[m] < target:
			l = m + 1
		case nums[m] > target:
			r = m - 1
		}
	}

	return false
}

// using sets/hash table
func intersection2(nums1 []int, nums2 []int) []int {
	set1 := make(map[int]struct{})
	set2 := make(map[int]struct{})

	for _, num := range nums1 {
		set1[num] = struct{}{}
	}
	for _, num := range nums2 {
		set2[num] = struct{}{}
	}

	var l int
	if len(set1) > len(set2) {
		l = len(set1)
	} else {
		l = len(set2)
	}

	ret := make([]int, 0, l)
	for num := range set1 {
		if _, ok := set2[num]; ok {
			ret = append(ret, num)
		}
	}

	return ret
}
