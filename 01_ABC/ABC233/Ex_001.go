package main

import (
	"bufio"
	"fmt"
	"os"
)

const M = 100005
const MX = M * 2

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	ps := make([][]int, MX)
	for i := 0; i < n; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		ps[x+y] = append(ps[x+y], x-y+M)
	}

	var q int
	fmt.Fscan(in, &q)
	a := make([]int, q)
	b := make([]int, q)
	k := make([]int, q)
	for i := 0; i < q; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y, &k[i])
		a[i] = x + y
		b[i] = x - y + M
	}
	wa := make([]int, q)
	for i := range wa {
		wa[i] = -1
	}
	ac := make([]int, q)
	for i := range ac {
		ac[i] = MX
	}

	for ti := 0; ti < 19; ti++ {
		num := make([]int, q)
		wj := make([]int, q)
		qs := make([][]int, MX)
		for i := 0; i < q; i++ {
			wj[i] = (wa[i] + ac[i]) / 2
			lx := a[i] - wj[i]
			rx := a[i] + wj[i] + 1
			lx = max(lx, 0)
			rx = min(rx, MX-1)
			qs[lx] = append(qs[lx], i)
			qs[rx] = append(qs[rx], q+i)
		}
		d := NewFenwick(MX)
		for x := 0; x < MX; x++ {
			for _, qi := range qs[x] {
				i := qi % q
				sign := 1
				if qi < q {
					sign = -1
				}
				ly := b[i] - wj[i]
				ry := b[i] + wj[i] + 1
				ly = max(ly, 0)
				ry = min(ry, MX)
				num[i] += d.Sum(ly, ry) * sign
			}
			for _, y := range ps[x] {
				d.Add(y, 1)
			}
		}
		for i := 0; i < q; i++ {
			if num[i] >= k[i] {
				ac[i] = wj[i]
			} else {
				wa[i] = wj[i]
			}
		}
	}

	for i := 0; i < q; i++ {
		fmt.Fprintln(out, ac[i])
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
