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

	const MOD = 998244353

	var q int
	fmt.Fscan(in, &q)
	X := make([]int, q)
	Y := make([]int, q)
	T := make([]int, q)
	for i := 0; i < q; i++ {
		var c string
		fmt.Fscan(in, &c, &X[i], &Y[i])
		if c == "+" {
			T[i] = 1
		} else {
			T[i] = -1
		}
		if X[i] < 0 {
			X[i] = -X[i]
			Y[i] = -Y[i]
		}
	}
	idx := make([]int, q)
	for i := range idx {
		idx[i] = i
	}
	sort.Slice(idx, func(i, j int) bool {
		return Y[idx[i]]*X[idx[j]] < X[idx[i]]*Y[idx[j]]
	})
	var rev [3 << 17]int
	for i := 0; i < q; i++ {
		rev[idx[i]] = i
	}
	A := NewFenwick(q)
	B := NewFenwick(q)
	ans := 0
	for i := 0; i < q; i++ {
		j := rev[i]
		BB := ((B.Sum(j, q) - B.Sum(0, j) + MOD) % MOD) * ((MOD + X[i]) % MOD) % MOD
		AA := ((A.Sum(j, q) - A.Sum(0, j) + MOD) % MOD) * ((MOD + Y[i]) % MOD) % MOD
		ans = (ans + ((((BB-AA+MOD)%MOD)*T[i] + MOD) % MOD)) % MOD
		fmt.Fprintln(out, ans)
		A.Add(j, (MOD+X[i]*T[i])%MOD)
		B.Add(j, (MOD+Y[i]*T[i])%MOD)
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

func (fen *Fenwick) lower_bound(x int) int {
	if fen.n == 0 {
		return 0
	}
	i := 0
	s := 0
	for k := 1 << log2(fen.n); k > 0; k >>= 1 {
		if i+k <= fen.n && s+int(fen.data[i+k-1]) < x {
			i += k
			s += int(fen.data[i-1])
		}
	}
	return i
}

func (fen *Fenwick) upper_bound(x int) int {
	if fen.n == 0 {
		return 0
	}
	i := 0
	s := 0
	for k := 1 << log2(fen.n); k > 0; k >>= 1 {
		if i+k <= fen.n && !(x < s+int(fen.data[i+k-1])) {
			i += k
			s += int(fen.data[i-1])
		}
	}
	return i
}

func log2(n int) int {
	var k int
	for k = 0; n != 0; n >>= 1 {
		k++
	}
	return k - 1
}
