/**
* 1971. Find if Path Exists in Graph (easy)
**/

package main

import "fmt"

// TODO: union find (disjoint set)
func validPath(n int, edges [][]int, source int, destination int) bool {
	root := map[int]int{}
	rank := make([]int, n)
	// for i := range rank {
	// 	rank[i] = 1
	// }

	var find func(x int) int
	find = func(x int) int {
		if x != root[x] {
			root[x] = find(root[x])
			fmt.Println("~ x:", x)
			fmt.Println("~ LEN ROOT:", len(root))
			fmt.Println("~ root:", root)
		}

		return root[x]
	}

	union := func(x, y int) {
		rootX := find(x)
		rootY := find(y)
		if rootX == rootY {
			return
		}

		if rank[rootX] < rank[rootY] {
			root[rootX] = rootY
		} else {
			root[rootY] = rootX

			if rank[rootX] == rank[rootY] {
				rank[rootX]++
			}
		}
	}

	for _, edge := range edges {
		union(edge[0], edge[1])
	}

	return find(source) == find(destination)
}

// bfs
func validPath2(n int, edges [][]int, source int, destination int) bool {
	// helpers
	destruct := func(s []int, vars ...*int) {
		for i, v := range s {
			*vars[i] = v
		}
	}

	al := make([][]int, n) // adjacency list
	for _, v := range edges {
		var x, y int
		destruct(v, &x, &y)
		al[x] = append(al[x], y)
		al[y] = append(al[y], x)
	}

	seen := map[int]bool{}
	queue := []int{source}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == destination {
			return true
		}

		for _, neighbor := range al[node] {
			if _, ok := seen[neighbor]; !ok {
				seen[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}

	return false
}

// dfs (TLE)
func validPath3(n int, edges [][]int, source int, destination int) bool {
	// helpers
	destruct := func(s []int, vars ...*int) {
		for i, v := range s {
			*vars[i] = v
		}
	}

	al := make([][]int, n) // adjacency list
	for _, v := range edges {
		var x, y int
		destruct(v, &x, &y)
		al[x] = append(al[x], y)
		al[y] = append(al[y], x)
	}

	seen := map[int]bool{}
	stack := []int{source}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		if node == destination {
			return true
		}

		for _, neighbor := range al[node] {
			if _, ok := seen[neighbor]; !ok {
				seen[neighbor] = true
				stack = append(stack, neighbor)
			}
		}
	}

	return false
}
