package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	type pair struct{ x, y int }
	e := make([][]pair, 1<<17)
	for i := 0; i < n; i++ {
		var m int
		fmt.Fscan(in, &m)
		p := make([]pair, m)
		mj := 0
		for j := 0; j < m; j++ {
			fmt.Fscan(in, &p[j].x, &p[j].y)
			if p[mj].x > p[j].x {
				mj = j
			} else if p[mj].x == p[j].x {
				if p[mj].y > p[j].y {
					mj = j
				}
			}
		}
		for j := 0; j < m; j += 2 {
			e[p[j].x] = append(e[p[j].x], pair{p[j+mj%2].y, p[j+1-mj%2].y})
		}
	}

	var q int
	fmt.Fscan(in, &q)

	res := [1 << 17]int{}
	Q := make([][]pair, 1<<17)
	for i := 0; i < q; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		Q[x] = append(Q[x], pair{y, i})
	}

	ft := NewFenwick(1 << 17)
	for x := 0; x < 1<<17; x++ {
		for _, p := range e[x] {
			ft.Add(p.x, 1)
			ft.Add(p.y, -1)
		}
		for _, p := range Q[x] {
			res[p.y] = ft.Sum(0, p.x+1)
		}
	}
	for i := 0; i < q; i++ {
		fmt.Fprintln(out, res[i])
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
