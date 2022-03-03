package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}

	fen := NewFenwick(n + 1)
	r := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		r[i] = -1
	}
	fen.Add(0, 1)
	for i := 0; i < n; i++ {
		fen.Add(i+1, fen.Sum(r[a[i]]+1, i+1))
		if r[a[i]] != -1 {
			fen.Add(r[a[i]]+1, -(fen.Sum(r[a[i]]+1, r[a[i]]+2) % mod))
		}
		r[a[i]] = i
	}
	fmt.Println(fen.Sum(1, n+1) % mod)
}

const mod = 998244353

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
	return int((fen.sum(r) - fen.sum(l) + mod) % mod)
}

func (fen *Fenwick) sum(r int) uint {
	s := uint(0)
	for r > 0 {
		s += fen.data[r-1]
		r -= r & -r
	}
	return s
}
