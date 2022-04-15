/**
* 66. Plus One (easy)
**/

package main

// addition with carry
func plusOne(digits []int) []int {
	n := len(digits)
	digits[n-1]++
	var carry int

	for i := n - 1; i >= 0; i-- {
		x := digits[i] + carry
		carry, digits[i] = x/10, x%10
	}

	if carry != 0 {
		return append([]int{carry}, digits...)
	} else {
		return digits
	}
}

// recursive
func plusOne2(digits []int) []int {
	n := len(digits)

	if n == 0 {
		return []int{1}
	}

	if digits[n-1] != 9 {
		return append(digits[:n-1], digits[n-1]+1)
	}

	return append(plusOne2(digits[:n-1]), 0)
}
