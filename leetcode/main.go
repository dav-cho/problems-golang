package main

import "fmt"

func main() {
	tests := [][]int{
		{1, 2, 3, 1},
		{1, 2, 3, 4},
		{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
	}
	answers := []bool{true, false, true}
	failed := 0

	for i := 0; i < len(tests); i++ {
		res := containsDuplicate(tests[i])
		ans := answers[i]
		if res != ans {
			failed++
		}
	}

	if failed == 0 {
		fmt.Println("OK")
	} else {
		fmt.Printf("Failed %d", failed)
	}
}
