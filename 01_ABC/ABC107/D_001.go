package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}

	ok, ng := 1, int(1e9+7)
	for ng-ok > 1 {
		mid := (ok + ng) / 2
		fen := NewFenwick(2*n + 1)
		fen.Add(n, 1)
		res, sum := 0, n
		for i := 0; i < n; i++ {
			if a[i] < mid {
				sum--
			} else {
				sum++
			}
			res += fen.Sum(0, sum+1)
			fen.Add(sum, 1)
		}
		if res >= (n*(n+1)/2+1)/2 {
			ok = mid
		} else {
			ng = mid
		}
	}
	fmt.Println(ok)
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
