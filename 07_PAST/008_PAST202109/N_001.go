package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(in, &N)

	A := make([]int, N)
	for i := range A {
		fmt.Fscan(in, &A[i])
	}

	B := make([]int, N)
	copy(B, A)
	sort.Ints(B)
	B = unique(B)

	for i := 0; i < N; i++ {
		A[i] = lowerBound(B, A[i])
	}

	const MOD = 1_000_000_007
	hi := NewFenwick(len(B))
	lo := NewFenwick(len(B))
	for i := 0; i < N; i++ {
		sumhi := hi.Sum(0, A[i]) % MOD
		sumlo := lo.Sum(A[i]+1, len(B)) % MOD
		hi.Add(A[i], 1+sumlo)
		lo.Add(A[i], 1+sumhi)
	}

	ans := 0
	ans += hi.Sum(0, len(B))
	ans += lo.Sum(0, len(B))
	ans += MOD - 2*N
	ans %= MOD
	fmt.Println(ans)
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

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}

func unique(a []int) []int {
	occurred := map[int]bool{}
	result := []int{}
	for i := range a {
		if occurred[a[i]] != true {
			occurred[a[i]] = true
			result = append(result, a[i])
		}
	}
	sort.Ints(result)
	n := len(result)
	for i := 0; i < n; i++ {
		a[i] = result[i]
	}
	return result
}
