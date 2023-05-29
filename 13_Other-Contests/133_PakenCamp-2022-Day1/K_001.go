package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	type pair struct {
		x, y int
	}

	var A [2 << 17]int
	var R [2 << 17][]pair

	var N int
	fmt.Fscan(in, &N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &A[i])
	}
	var Q int
	fmt.Fscan(in, &Q)
	for i := 0; i < Q; i++ {
		var l, r int
		fmt.Fscan(in, &l, &r)
		R[l-1] = append(R[l-1], pair{r, i})
	}
	S := make([]int, 0)
	BIT := NewFenwick(N)
	ans := make([]int, Q)
	for i := N - 1; i >= 0; i-- {
		for len(S) != 0 && A[S[len(S)-1]] <= A[i] {
			BIT.Add(S[len(S)-1], -1)
			S = S[:len(S)-1]
		}
		BIT.Add(i, 1)
		S = append(S, i)
		for _, q := range R[i] {
			ans[q.y] = BIT.Sum(i, q.x)
		}
	}
	for i := 0; i < Q; i++ {
		fmt.Println(ans[i])
	}
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
