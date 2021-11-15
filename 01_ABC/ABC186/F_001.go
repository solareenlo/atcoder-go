package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var h, w, m int
	fmt.Fscan(in, &h, &w, &m)

	X := make([]int, w)
	for i := range X {
		X[i] = h
	}
	Y := make([]int, h)
	for i := range Y {
		Y[i] = w
	}

	for i := 0; i < m; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		x--
		y--
		if X[y] > x {
			X[y] = x
		}
		if Y[x] > y {
			Y[x] = y
		}
	}

	res := 0
	B := make([][]int, w+1)
	F := NewFenwick(h + 1)
	for i := 0; i < max(X[0], Y[0]); i++ {
		if i < Y[0] {
			res += X[i]
		}
		if i < X[0] {
			res += Y[i]
			F.Add(i, 1)
			B[Y[i]] = append(B[Y[i]], i)
		}
	}

	for i := 0; i < Y[0]; i++ {
		for _, x := range B[i] {
			F.Add(x, -1)
		}
		res -= F.Sum(0, X[i])
	}

	fmt.Println(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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
