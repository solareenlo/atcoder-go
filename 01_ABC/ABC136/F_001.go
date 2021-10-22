package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type pair struct{ x, y int }

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	mod := 998244353

	p2 := make([]int, n+1)
	for i := range p2 {
		p2[i] = 1
	}
	for i := 1; i < n+1; i++ {
		p2[i] = p2[i-1] * 2 % mod
	}

	xy := make([]pair, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &xy[i].x, &xy[i].y)
	}
	sort.Slice(xy, func(i, j int) bool {
		return xy[i].y < xy[j].y
	})

	for i := 0; i < n; i++ {
		xy[i].y = i
	}
	sort.Slice(xy, func(i, j int) bool {
		return xy[i].x < xy[j].x
	})

	ft := NewFenwick(n)
	res := 0
	for i := 0; i < n; i++ {
		si := p2[n] - 1
		c00 := ft.Sum(0, xy[i].y)
		c01 := xy[i].y - c00
		c10 := i - c00
		c11 := n - 1 - c00 - c01 - c10
		si = (si - p2[c00+c01] - p2[c10+c11] - p2[c00+c10] - p2[c01+c11] + mod*4) % mod
		si = (si + p2[c00] + p2[c01] + p2[c10] + p2[c11]) % mod
		res = (res + si) % mod
		ft.Add(xy[i].y, 1)
	}
	fmt.Println(res)
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
