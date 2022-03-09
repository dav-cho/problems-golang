package main

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func main() {
	// test1 := coinChange1([]int{1, 2, 5}, 11)
	// println(test1)
	// test2 := coinChange1([]int{2}, 3)
	// println(test2)

	test1 := findCircleNum([][]int{
		{1, 1, 0},
		{1, 1, 0},
		{0, 0, 1},
	})
	println(test1)
}
