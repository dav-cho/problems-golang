/**
* 189. Rotate Array (medium)
**/

package main

// cyclic replacements
func rotate(nums []int, k int) {
	n := len(nums)
	sIdx, count := 0, n

	for count > 0 {
		prev := nums[sIdx]
		cIdx := sIdx
		for true {
			nIdx := (cIdx + k) % n
			prev, nums[nIdx] = nums[nIdx], prev
			cIdx = nIdx
			count--

			if cIdx == sIdx {
				break
			}
		}

		sIdx++
	}
}

// extra array
func rotate2(nums []int, k int) {
	n := len(nums)
	rot_arr := make([]int, n)
	for i := 0; i < n; i++ {
		rot_arr[(i+k)%n] = nums[i]
	}

	copy(nums, rot_arr)
}

// reverse
func rotate3(nums []int, k int) {
	n := len(nums)
	k = k % n

	move(&nums, 0, n-1)
	move(&nums, 0, k-1)
	move(&nums, k, n-1)
}

func move(nums *[]int, start, end int) {
	for start < end {
		(*nums)[start], (*nums)[end] = (*nums)[end], (*nums)[start]
		start++
		end--
	}
}
