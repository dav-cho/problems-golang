/**
* 323. Number of Connected Components in an Undirected Graph (medium)
**/

package main

// unionfind / disjoint set
func countComponents(n int, edges [][]int) int {
	root, rank, count := make([]int, n), make([]int, n), n
	for i := range root {
		root[i] = i
		rank[i] = 0
	}
	uf := UnionFind3{count, root, rank}
	for _, e := range edges {
		uf.union(e[0], e[1])
	}

	return uf.count
}

type UnionFind3 struct {
	count      int
	root, rank []int
}

func (uf *UnionFind3) find(x int) int {
	if x == uf.root[x] {
		return x
	}

	uf.root[x] = uf.find(uf.root[x])
	return uf.root[x]
}

func (uf *UnionFind3) union(x, y int) {
	rx, ry := uf.find(x), uf.find(y)
	if rx == ry {
		return
	}

	if uf.rank[rx] > uf.rank[ry] {
		uf.root[ry] = rx
	} else {
		uf.root[rx] = ry
		if uf.rank[rx] == uf.rank[ry] {
			uf.rank[rx]++
		}
	}

	uf.count--
}
