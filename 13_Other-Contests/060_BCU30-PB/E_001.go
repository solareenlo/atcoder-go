package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n, k, q int
	fmt.Fscan(in, &n, &k, &q)

	s := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &s[i])
	}

	a := make([][]int, n+1)
	for i := range a {
		a[i] = make([]int, n+1)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			a[i+1][j+1] = a[i+1][j] + a[i][j+1] - a[i][j];
			if s[i][j] == '#' {
				a[i+1][j+1]++
			}
		}
	}

	m := n - k + 1
	ok := make([][]int, m)
	for i := range ok {
		ok[i] = make([]int, m)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			if a[i+k][j+k] - a[i+k][j] - a[i][j+k] + a[i][j] == 0 {
				ok[i][j] = 1
			}
		}
	}

	uf := New(m*m)
	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			if ok[i][j] != 0 {
				if i < m-1 && ok[i+1][j] != 0 {
					uf.Merge(i*m+j, (i+1)*m+j)
				}
				if j < m-1 && ok[i][j+1] != 0 {
					uf.Merge(i*m+j, i*m+(j+1))
				}
			}
		}
	}

	for i := 0; i < q; i++ {
		var r1, c1, r2, c2 int
		fmt.Fscan(in, &r1, &c1, &r2, &c2)
		r1--
		c1--
		r2--
		c2--
		if uf.Same(r1*m+c1, r2*m+c2) == true {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}

type UnionFind struct {
	root []int
	size []int
}

func New(size int) *UnionFind {
	return new(UnionFind).init(size)
}

func (uf *UnionFind) init(size int) *UnionFind {
	uf = new(UnionFind)
	uf.root = make([]int, size)
	uf.size = make([]int, size)

	for i := 0; i < size; i++ {
		uf.root[i] = i
		uf.size[i] = 1
	}

	return uf
}

func (uf *UnionFind) Merge(p, q int) {
	qRoot := uf.Root(q)
	pRoot := uf.Root(p)

	if uf.size[qRoot] < uf.size[pRoot] {
		uf.root[qRoot] = uf.root[pRoot]
		uf.size[pRoot] += uf.size[qRoot]
	} else {
		uf.root[pRoot] = uf.root[qRoot]
		uf.size[qRoot] += uf.size[pRoot]
	}
}

func (uf *UnionFind) Root(p int) int {
	if p > len(uf.root)-1 {
		return -1
	}

	for uf.root[p] != p {
		uf.root[p] = uf.root[uf.root[p]]
		p = uf.root[p]
	}

	return p
}

func (uf *UnionFind) Same(p, q int) bool {
	return uf.Root(p) == uf.Root(q)
}
