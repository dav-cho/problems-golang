/**
* 136. Single Number (easy)
**/

package main

// bitwise xor
func singleNumber(nums []int) int {
	ret := 0

	for _, num := range nums {
		ret ^= num
	}

	return ret
}

// hash table
func singleNumber2(nums []int) int {
	seen := make(map[int]struct{})
	a, b := 0, 0

	for _, num := range nums {
		seen[num] = struct{}{}
		a += num
	}

	for num := range seen {
		b += num
	}

	return 2*b - a
}
