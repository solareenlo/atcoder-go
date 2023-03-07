package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(in, &N)
	V := make([]int, 26)
	uf := NewUnionFind()
	for i := 0; i < N; i++ {
		var w string
		fmt.Fscan(in, &w)
		a := int(w[0] - 'a')
		b := int(w[len(w)-1] - 'a')
		V[a]++
		V[b]++
		uf.Merge(a, b)
	}
	T := make([]int, 26)
	for i := 0; i < 26; i++ {
		T[uf.Root(i)] += V[i] % 2
	}
	z := -1
	for i := 0; i < 26; i++ {
		if V[i] != 0 && uf.Root(i) == i {
			if T[i]/2 != 0 {
				z += T[i] / 2
			} else {
				z++
			}
		}
	}
	fmt.Println(z)
}

type UnionFind struct {
	root []int
	size []int
}

func NewUnionFind() *UnionFind {
	uf := new(UnionFind)
	uf.root = make([]int, 0)
	uf.size = make([]int, 0)
	return uf
}

func (uf *UnionFind) extend(N int) {
	for i := len(uf.root); i <= N; i++ {
		uf.root = append(uf.root, i)
		uf.size = append(uf.size, 1)
	}
}

func (uf *UnionFind) Root(i int) int {
	uf.extend(i)
	if uf.root[i] == i {
		uf.root[i] = i
	} else {
		uf.root[i] = uf.Root(uf.root[i])
	}
	return uf.root[i]
}

func (uf *UnionFind) Merge(a, b int) {
	uf.extend(a)
	uf.extend(b)
	a = uf.Root(a)
	b = uf.Root(b)
	if a != b {
		if uf.size[a] > uf.size[b] {
			a, b = b, a
		}
		uf.size[b] += uf.size[a]
		uf.root[a] = b
	}
}

func (uf UnionFind) Same(a, b int) bool {
	uf.extend(a)
	uf.extend(b)
	return uf.Root(a) == uf.Root(b)
}

func (uf UnionFind) Size(x int) int {
	uf.extend(x)
	return uf.size[uf.Root(x)]
}
