package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	type c struct{ a, b int }
	d := make([]c, n)
	m := map[int]int{}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &d[i].a)
		m[d[i].a] = 0
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &d[i].b)
	}

	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	k := 0
	for i := range keys {
		m[keys[i]] = k
		k++
	}
	sort.Slice(d, func(i, j int) bool {
		return (d[i].b == d[j].b && d[i].a > d[j].a) || d[i].b < d[j].b
	})

	t := NewFenwick(n)
	res := 0
	for i := 0; i < n; i++ {
		if i != 0 && d[i].a == d[i-1].a && d[i].b == d[i-1].b {
			k++
		} else {
			k = 0
		}
		p := m[d[i].a]
		t.Add(p, 1)
		res += k + t.Sum(p, n)
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
