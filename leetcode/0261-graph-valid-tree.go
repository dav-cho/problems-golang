/**
* 261. Graph Valid Tree (medium)
**/

package main

// unionfind / disjoint set
func validTree(n int, edges [][]int) bool {
	if len(edges) != n-1 {
		return false
	}

	root, rank := make([]int, n), make([]int, n)
	for i := range root {
		root[i] = i
		rank[i] = 0
	}
	uf := UnionFind2{root, rank}
	for _, e := range edges {
		if !uf.union(e[0], e[1]) {
			return false
		}
	}

	return true
}

type UnionFind2 struct {
	root, rank []int
}

func (uf *UnionFind2) find(x int) int {
	if x == uf.root[x] {
		return x
	}

	uf.root[x] = uf.find(uf.root[x])
	return uf.root[x]
}

func (uf *UnionFind2) union(x, y int) bool {
	rx, ry := uf.find(x), uf.find(y)
	if rx == ry {
		return false
	}

	if uf.rank[rx] > uf.rank[ry] {
		uf.root[ry] = rx
	} else {
		uf.root[rx] = ry
		if uf.rank[rx] == uf.rank[ry] {
			uf.rank[rx]++
		}
	}

	return true
}
