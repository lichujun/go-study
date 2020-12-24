package algorithm

import (
	"fmt"
	"testing"
)

func TestUnionFind(t *testing.T) {
	uf := NewArrayUnionFind(10)
	uf.Union(1, 2)
	uf.Union(2, 5)
	uf.Union(3, 4)
	uf.Union(1, 3)
	fmt.Println(uf.Count())
}

type ArrayUnionFind struct {
	parent []int
	size   []int
	count  int
}

type UnionFind interface {
	Count() int
	Union(p int, q int)
	Connected(p int, q int) bool
}

func NewArrayUnionFind(n int) UnionFind {
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}
	return &ArrayUnionFind{
		parent: parent,
		size:   size,
		count:  n,
	}
}

func (auf *ArrayUnionFind) Count() int {
	return auf.count
}

func (auf *ArrayUnionFind) Union(p int, q int) {
	pRoot := auf.find(p)
	qRoot := auf.find(q)
	if pRoot == qRoot {
		return
	}
	if auf.size[pRoot] > auf.size[qRoot] {
		auf.parent[qRoot] = pRoot
		auf.size[pRoot] += auf.size[qRoot]
	} else {
		auf.parent[pRoot] = qRoot
		auf.size[qRoot] += auf.size[pRoot]
	}
	auf.count--
}

func (auf *ArrayUnionFind) Connected(p int, q int) bool {
	pRoot := auf.find(p)
	qRoot := auf.find(q)
	return pRoot == qRoot
}

func (auf *ArrayUnionFind) find(p int) int {
	for auf.parent[p] != p {
		auf.parent[p] = auf.parent[auf.parent[p]]
		p = auf.parent[p]
	}
	return p
}
