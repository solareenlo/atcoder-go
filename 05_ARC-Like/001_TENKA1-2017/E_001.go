package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var n int

func get(A, B, C []float64) float64 {
	type pair struct{ x, y float64 }
	p := make([]pair, 0)
	for i := 0; i < n; i++ {
		p = append(p, pair{-A[i] / B[i], C[i] / B[i]})
	}
	sort.Slice(p, func(i, j int) bool {
		return p[i].x < p[j].x
	})
	ok := -1e9
	ng := 1e9
	for j := 0; j < 70; j++ {
		mid := (ok + ng) / 2.0
		type pair2 struct {
			f float64
			i int
		}
		pp := make([]pair2, n)
		for i := 0; i < n; i++ {
			pp[i] = pair2{mid*p[i].x + p[i].y, i}
		}
		sort.Slice(pp, func(i, j int) bool {
			return pp[i].f < pp[j].f
		})
		F := NewFenwick(n)
		sum := 0
		for i := 0; i < n; i++ {
			sum += F.Sum(0, pp[i].i)
			F.Add(pp[i].i, 1)
		}
		if sum < ((n*(n-1)/2)+1)/2 {
			ok = mid
		} else {
			ng = mid
		}
	}

	return ok
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n)
	a := make([]float64, n)
	b := make([]float64, n)
	c := make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i], &b[i], &c[i])
	}

	fmt.Println(get(a, b, c), get(b, a, c))
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
