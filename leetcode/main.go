package main

import "fmt"

func main() {
	tests := [][]interface{}{
		{[]int{2, 7, 11, 15}, 9},
		{[]int{3, 2, 4}, 6},
		{[]int{3, 3}, 6},
	}
	answers := [][]int{
		{0, 1},
		{1, 2},
		{0, 1},
	}
	failed := 0

	for i := 0; i < len(tests); i++ {
		res := twoSum(tests[i][0].([]int), tests[i][1].(int))
		ans := answers[i]

		for i := 0; i < len(tests[i]); i++ {
			if res[i] != ans[i] {
				failed++
			}
		}
	}

	if failed == 0 {
		fmt.Println("OK")
	} else {
		fmt.Printf("Failed %d\n", failed)
	}
}
