package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const MOD = 1000000007

	var n, k int
	fmt.Fscan(in, &n, &k)

	I := make([]int, n)
	for i := range I {
		I[i] = i
	}

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	sort.Slice(I, func(i, j int) bool {
		return a[I[i]] < a[I[j]]
	})

	p := 0
	for i := 0; i < n; i++ {
		if a[I[i]] > p {
			p = a[I[i]]
			a[I[i]] = i
		} else {
			a[I[i]] = a[I[i-1]]
		}
	}

	b := NewFenwick(n)
	B := make([]int, n+1)
	t := 0
	p = 0
	B[1] = 1
	for i := 0; i < n; i++ {
		if i != 0 {
			B[i+1] = (B[i]*2%MOD - B[p] + MOD) % MOD
		}
		t += b.Sum(a[i]+1, n)
		b.Add(a[i], 1)
		for ; t > k; p++ {
			b.Add(a[p], -1)
			t -= b.Sum(0, a[p])
		}
	}
	fmt.Println((B[n] - B[p] + MOD) % MOD)
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
