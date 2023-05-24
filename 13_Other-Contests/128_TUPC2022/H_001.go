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
	p := make([]int, n)
	q := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &p[i])
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &q[i])
	}
	pi := make([]int, n)
	qi := make([]int, n)
	pf := NewFenwick(n)
	qf := NewFenwick(n)
	pqf := NewFenwick(n)
	for i := 0; i < n; i++ {
		pi[i] = p[i] - 1 - pf.Sum(0, p[i])
		pf.Add(p[i]-1, 1)
	}
	for i := 0; i < n; i++ {
		qi[i] = q[i] - 1 - qf.Sum(0, q[i])
		qf.Add(q[i]-1, 1)
	}
	pqi := make([]int, n)
	carry := false
	for i := n - 1; i >= 0; i-- {
		pqi[i] = qi[i] - pi[i]
		if carry {
			pqi[i]--
		}
		carry = false
		if pqi[i] < 0 {
			carry = true
			pqi[i] += n - i
		}
	}

	for i := 0; i < n; i++ {
		l := 0
		r := n
		for r-l > 1 {
			x := (r + l) / 2
			if x-pqf.Sum(0, x+1) > pqi[i] || (x-pqf.Sum(0, x+1) == pqi[i] && pqf.Sum(x, x+1) == 1) {
				r = x
			} else {
				l = x
			}
		}
		pqf.Add(l, 1)
		if i+1 >= n {
			fmt.Println(l + 1)
		} else {
			fmt.Printf("%d ", l+1)
		}
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
