package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, q int
	fmt.Fscan(in, &n, &q)
	Y := make([]int, 0)
	type tuple struct{ x, t, y, c int }
	queries := make([]tuple, 0)
	for i := 0; i < n; i++ {
		var x, y, d, c int
		fmt.Fscan(in, &x, &y, &d, &c)
		Y = append(Y, y)
		Y = append(Y, y+d+1)
		queries = append(queries, tuple{x, 1, y, c})
		queries = append(queries, tuple{x, 1, y + d + 1, -c})
		queries = append(queries, tuple{x + d + 1, 1, y, -c})
		queries = append(queries, tuple{x + d + 1, 1, y + d + 1, c})
	}
	for i := 0; i < q; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		Y = append(Y, b)
		queries = append(queries, tuple{a, 2, b, i})
	}
	sort.Ints(Y)
	sort.Slice(queries, func(i, j int) bool {
		return queries[i].x < queries[j].x || (queries[i].x == queries[j].x && queries[i].t < queries[j].t)
	})

	ret := make([]int, q)
	fw := NewFenwick(2*n + q)
	for i := range queries {
		queries[i].y = lowerBound(Y, queries[i].y)
		if queries[i].t == 1 {
			fw.Add(queries[i].y, queries[i].c)
		} else {
			ret[queries[i].c] = fw.Sum(0, queries[i].y+1)
		}
	}
	for i := 0; i < q; i++ {
		fmt.Fprintln(out, ret[i])
	}
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}

type Fenwick struct {
	n    int
	data []uint
}

func NewFenwick(n int) *Fenwick {
	fen := &Fenwick{
		n:    n,
		data: make([]uint, n),
	}
	for idx := range fen.data {
		fen.data[idx] = 0
	}
	return fen
}

func (fen *Fenwick) Add(pos, x int) {
	if !(0 <= pos && pos < fen.n) {
		panic("")
	}
	pos++
	for pos <= fen.n {
		fen.data[pos-1] += uint(x)
		pos += pos & -pos
	}
}

func (fen *Fenwick) Sum(l, r int) int {
	if !(0 <= l && l <= r && r <= fen.n) {
		panic("")
	}
	return int(fen.sum(r) - fen.sum(l))
}

func (fen *Fenwick) sum(r int) uint {
	s := uint(0)
	for r > 0 {
		s += fen.data[r-1]
		r -= r & -r
	}
	return s
}
