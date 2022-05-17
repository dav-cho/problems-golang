/**
* 547. Number of Provinces (medium)
**/

package main

// disjoint set (union find)
func findCircleNum(isConnected [][]int) int {
	l := len(isConnected)
	uf := newUnionFind(l)
	for i := range isConnected {
		for j := i; i < l; i++ {
			if isConnected[i][j] == 1 {
				uf.union(i, j)
			}
		}
	}

	return uf.count
}

type UnionFind struct {
	root, rank []int
	count      int
}

func newUnionFind(size int) *UnionFind {
	root, rank := make([]int, size), make([]int, size)
	for i := range rank {
		root[i] = i
	}

	return &UnionFind{root, rank, size}
}

func (uf *UnionFind) find(x int) int {
	if x == uf.root[x] {
		return x
	}

	uf.root[x] = uf.find(uf.root[x])
	return uf.root[x]
}

func (uf *UnionFind) union(x, y int) {
	rootX := uf.find(x)
	rootY := uf.find(y)
	if rootX == rootY {
		return
	}

	if uf.rank[rootX] > uf.rank[rootY] {
		uf.root[rootY] = rootX
	} else {
		uf.root[rootX] = rootY

		if uf.rank[rootX] == uf.rank[rootY] {
			uf.rank[rootX]++
		}
	}

	uf.count--
}
