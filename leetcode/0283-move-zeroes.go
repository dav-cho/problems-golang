/**
* 283. Move Zeroes (easy)
**/

package main

// most optimal - one pass
func moveZeroes(nums []int)  {
    var k int

    for i := 0; i < len(nums); i++ {
        if nums[i] != 0 {
            nums[i], nums[k] = nums[k], nums[i]
            k++
        }
    }
}

// space optimal - loop twice
func moveZeroes2(nums []int)  {
    var k int

    for _, num := range nums {
        if num != 0 {
            nums[k] = num
            k += 1
        }
    }

    for i := k; i < len(nums); i++ {
        nums[i] = 0
    }
}
