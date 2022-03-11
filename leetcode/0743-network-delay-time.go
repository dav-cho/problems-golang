/**
* 743. Network Delay Time (medium)
**/

package main

import (
	"math"
	"sort"
)

// dfs
func networkDelayTime1(times [][]int, n int, k int) int {
	// helpers
	destruct := func(s []int, vars ...*int) {
		for i, v := range s {
			*vars[i] = v
		}
	}

	max := math.MaxInt32
	g := map[int][][]int{}
	for _, s := range times {
		var u, v, w int // u: node, v: neighbor, w: time
		destruct(s, &u, &w, &v)
		g[u] = append(g[u], []int{v, w})
	}

	d := map[int]int{}
	for i := 1; i <= n; i++ {
		d[i] = max
	}

	var dfs func(n int, e int)
	dfs = func(n int, e int) { // n: node, e: elapsed time
		if e >= d[n] {
			return
		}

		d[n] = e
		sort.SliceStable(g[n], func(i, j int) bool { return g[n][i][0] < g[n][j][0] })
		for _, val := range g[n] {
			var w, v int // w: time, v: neighbor
			destruct(val, &w, &v)
			dfs(v, e+w)
		}
	}

	dfs(k, 0)

	ret := -max
	for _, v := range d {
		if v > ret {
			ret = v
		}
	}

	if ret < max {
		return ret
	}

	return -1
}

// TODO: bfs
// func networkDelayTime(times [][]int, n int, k int) int {

// }

// dijkstra's algorithm
// func networkDelayTime3(times [][]int, n int, k int) int {
func networkDelayTime(times [][]int, n int, k int) int {
	// helpers
	min := func(x, y int) int {
		if x < y {
			return x
		}

		return y
	}
	destruct := func(s []int, vars ...*int) {
		for i, v := range s {
			*vars[i] = v
		}
	}
	mapVals := func(m map[int]int) []int {
		vals := make([]int, 0, len(m))
		for _, v := range m {
			vals = append(vals, v)
		}

		return vals
	}
	maxSlice := func(s []int) int {
		ret := -math.MaxInt32

		for _, v := range s {
			if v > ret {
				ret = v
			}
		}

		return ret
	}

	max := math.MaxInt32
	seen := map[int]bool{}
	g := make(map[int][][]int) // graph
	for _, val := range times {
		var u, v, w int
		destruct(val, &u, &v, &w)
		g[u] = append(g[u], []int{v, w})
	}

	d := map[int]int{} // distance
	for node := 1; node <= n; node++ {
		d[node] = max
	}
	d[k] = 0

	for {
		cn := -1  // candidate node
		cd := max // candidate dist
		for node := 1; node <= n; node++ {
			if _, ok := seen[node]; !ok && d[node] < cd {
				cn = node
				cd = d[node]
			}
		}
		if cn < 0 {
			break
		}
		seen[cn] = true

		for _, v := range g[cn] {
			var nb, t int // nb = neighbor, t = time
			destruct(v, &nb, &t)
			d[nb] = min(d[nb], d[cn]+t)
		}
	}

	ret := maxSlice(mapVals(d))
	if ret < max {
		return ret
	}

	return -1
}
