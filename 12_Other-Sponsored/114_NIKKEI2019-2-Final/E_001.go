package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type interval struct {
	l, r, id int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m, q int
	fmt.Fscan(in, &n, &m, &q)
	var S string
	fmt.Fscan(in, &S)
	s := strings.Split(S, "")
	I := make([]interval, m)
	for i := 0; i < m; i++ {
		var l, r int
		fmt.Fscan(in, &l, &r)
		I[i] = interval{l: l - 1, r: r, id: i}
	}

	ans := make([]int, q+1)
	F := NewFenwick(n)
	for i := 0; i < n; i++ {
		if s[i] == "1" {
			F.Add(i, 1)
		}
	}
	for i := 0; i < m; i++ {
		if F.Sum(I[i].l, I[i].r) > 0 {
			ans[0]++
		}
	}

	J1 := make([]interval, q)
	J2 := make([]interval, q)
	J3 := make([]interval, q)
	pm := make([]int, q)
	for i := 0; i < q; i++ {
		var x int
		fmt.Fscan(in, &x)
		x--
		if s[x] == "0" {
			pm[i] = 1
		} else {
			pm[i] = -1
		}
		if pm[i] == -1 {
			s[x] = "0"
			F.Add(x, -1)
		}
		l := F.lower_bound(F.Sum(0, x))
		if s[l] == "1" {
			l++
		}
		r := F.upper_bound(F.Sum(0, x))
		J1[i] = interval{l, r, i}
		J2[i] = interval{l, x, i}
		J3[i] = interval{x + 1, r, i}
		if pm[i] == 1 {
			s[x] = "1"
			F.Add(x, 1)
		}
	}

	K := make([]Pair, 0)
	for i := 0; i < len(I); i++ {
		K = append(K, Pair{I[i], -1})
	}
	for i := 0; i < len(J1); i++ {
		K = append(K, Pair{J1[i], 1})
	}
	for i := 0; i < len(J2); i++ {
		K = append(K, Pair{J2[i], 2})
	}
	for i := 0; i < len(J3); i++ {
		K = append(K, Pair{J3[i], 3})
	}
	sort.Slice(K, func(i, j int) bool {
		if K[i].x.r == K[j].x.r {
			if K[i].x.l == K[j].x.l {
				return K[i].y < K[j].y
			}
			return -K[i].x.l < -K[j].x.l
		}
		return K[i].x.r < K[j].x.r
	})

	Left := NewFenwick(n)
	for _, p := range K {
		k := p.x
		if p.y < 0 {
			Left.Add(k.l, 1)
		} else {
			var tmp int
			if p.y == 1 {
				tmp = 1
			} else {
				tmp = -1
			}
			ans[k.id+1] += pm[k.id] * tmp * Left.Sum(k.l, k.r)
		}
	}

	for i := 0; i < q; i++ {
		ans[i+1] += ans[i]
		fmt.Fprintln(out, ans[i+1])
	}
}

type Pair struct {
	x interval
	y int
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
