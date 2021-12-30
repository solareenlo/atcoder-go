package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const N = 500050

var (
	sz, res int
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m, q int
	fmt.Fscan(in, &n, &m, &q)

	t := make([]int, q)
	X := make([]int, q)
	Y := make([]int, q)
	my := map[int]bool{}
	my[0] = true
	yy := make([]int, 0)
	for i := 0; i < q; i++ {
		fmt.Fscan(in, &t[i], &X[i], &Y[i])
		X[i]--
		my[Y[i]] = true
	}
	for k := range my {
		yy = append(yy, k)
	}
	sort.Ints(yy)

	ym := map[int]int{}
	for i := 0; i < len(yy); i++ {
		ym[yy[i]] = i
	}

	A := make([]int, n)
	B := make([]int, m)
	A1 := NewFenwick(len(yy))
	B1 := NewFenwick(len(yy))
	A2 := NewFenwick(len(yy))
	B2 := NewFenwick(len(yy))
	sum := 0
	A2.Add(0, n)
	B2.Add(0, m)
	for i := 0; i < q; i++ {
		x := X[i]
		y := Y[i]
		if t[i] == 1 {
			sum -= B2.Sum(0, ym[A[x]]) * A[x]
			sum += B2.Sum(0, ym[y]) * y
			sum += B1.Sum(0, ym[A[x]])
			sum -= B1.Sum(0, ym[y])
			A1.Add(ym[A[x]], -A[x])
			A1.Add(ym[y], y)
			A2.Add(ym[A[x]], -1)
			A2.Add(ym[y], 1)
			A[x] = y
		} else {
			sum -= A2.Sum(0, ym[B[x]]) * B[x]
			sum += A2.Sum(0, ym[y]) * y
			sum += A1.Sum(0, ym[B[x]])
			sum -= A1.Sum(0, ym[y])
			B1.Add(ym[B[x]], -B[x])
			B1.Add(ym[y], y)
			B2.Add(ym[B[x]], -1)
			B2.Add(ym[y], 1)
			B[x] = y
		}
		fmt.Fprintln(out, sum)
	}
}

type Fenwick struct {
	n    int
	data []uint
}

func NewFenwick(n int) *Fenwick {
	fen := &Fenwick{
		n:    n,
		data: make([]uint, n+1),
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
