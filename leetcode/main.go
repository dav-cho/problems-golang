package main

func main() {
	test1 := networkDelayTime([][]int{
		{2, 1, 1},
		{2, 3, 1},
		{3, 4, 1},
	}, 4, 2)
	println(test1)
}
