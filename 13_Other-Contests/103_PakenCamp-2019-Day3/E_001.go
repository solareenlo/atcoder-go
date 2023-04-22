package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	type pair struct {
		x, y int
	}

	type tuple struct {
		x, y, z int
	}

	var dx []int = []int{1, 0, -1, 0}
	var dy []int = []int{0, -1, 0, 1}

	var h, w int
	fmt.Fscan(in, &h, &w)
	var B [1500][]string
	for i := 0; i < h; i++ {
		var t string
		fmt.Fscan(in, &t)
		B[i] = strings.Split(t, "")
	}

	var dp [1500][1500]int
	for i := 0; i < h; i++ {
		if B[i][w-1] == "." {
			dp[i][w-1] = 1
		}
	}
	for j := 0; j < w; j++ {
		if B[h-1][j] == "." {
			dp[h-1][j] = 1
		}
	}

	for i := h - 2; i >= 0; i-- {
		for j := w - 2; j >= 0; j-- {
			if B[i][j] == "." {
				dp[i][j] = min(dp[i+1][j], dp[i][j+1], dp[i+1][j+1]) + 1
			}
		}
	}

	var Ev1 [1501][]pair
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if B[i][j] == "." {
				Ev1[dp[i][j]] = append(Ev1[dp[i][j]], pair{i, j})
			}
		}
	}

	var Ev2 [1501][]tuple
	var q int
	fmt.Fscan(in, &q)
	for i := 0; i < q; i++ {
		var x, y, l int
		fmt.Fscan(in, &y, &x, &l)
		x--
		y--
		Ev2[l] = append(Ev2[l], tuple{i, y, x})
	}

	ans := make([]int, q)
	U := NewUnionFind(h * w)
	var ok [1500][1500]bool
	for sz := 1500; sz > 0; sz-- {
		for _, e := range Ev1[sz] {
			y := e.x
			x := e.y
			ok[y][x] = true
			for k := 0; k < 4; k++ {
				xx := x + dx[k]
				yy := y + dy[k]
				if 0 <= xx && xx < w && 0 <= yy && yy < h && ok[yy][xx] {
					U.Merge(y*w+x, yy*w+xx)
				}
			}
		}
		for _, e := range Ev2[sz] {
			id := e.x
			y := e.y
			x := e.z
			ans[id] = U.Size(y*w + x)
		}
	}

	for i := 0; i < q; i++ {
		fmt.Fprintln(out, ans[i])
	}
}

type UnionFind struct {
	root []int
	size []int
}

func NewUnionFind(size int) *UnionFind {
	uf := new(UnionFind)
	uf.root = make([]int, size)
	uf.size = make([]int, size)

	for i := 0; i < size; i++ {
		uf.root[i] = i
		uf.size[i] = 1
	}

	return uf
}

func (uf *UnionFind) Merge(p, q int) bool {
	q = uf.Root(q)
	p = uf.Root(p)

	if q == p {
		return false
	}

	if uf.size[q] < uf.size[p] {
		uf.root[q] = uf.root[p]
		uf.size[p] += uf.size[q]
	} else {
		uf.root[p] = uf.root[q]
		uf.size[q] += uf.size[p]
	}
	return true
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

func (uf *UnionFind) Size(x int) int {
	return uf.size[uf.Root(x)]
}

func min(a ...int) int {
	res := a[0]
	for i := range a {
		if res > a[i] {
			res = a[i]
		}
	}
	return res
}
