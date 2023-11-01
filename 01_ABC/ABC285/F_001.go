package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var tr [26]*Fenwick
	var n, q int
	var tmp string
	fmt.Fscan(in, &n, &tmp, &q)
	tmp = " " + tmp
	a := strings.Split(tmp, "")
	for i := 0; i < 26; i++ {
		tr[i] = NewFenwick(n + 1)
	}
	for i := 1; i <= n; i++ {
		tr[a[i][0]-97].Add(i, 1)
	}
	var s, t [26]int
	for q > 0 {
		q--
		var op, x int
		fmt.Fscan(in, &op, &x)
		if (op & 1) != 0 {
			var c string
			fmt.Fscan(in, &c)
			tr[a[x][0]-97].Add(x, -1)
			a[x] = c
			tr[a[x][0]-97].Add(x, 1)
		} else {
			var y int
			fmt.Fscan(in, &y)
			for i := 0; i < 26; i++ {
				s[i] = tr[i].Sum(x, y+1)
				t[i] = tr[i].Sum(1, n+1)
			}
			pl, pr := 0, 25
			flg := true
			for s[pl] == 0 {
				pl++
			}
			for s[pr] == 0 {
				pr--
			}
			for i := pl + 1; i < pr; i++ {
				if flg && s[i] == t[i] {
					flg = true
				} else {
					flg = false
				}
			}
			for i := 0; i < 26; x, i = x+s[i], i+1 {
				if flg && tr[i].Sum(x, x+s[i]) == s[i] {
					flg = true
				} else {
					flg = false
				}
			}
			if flg {
				fmt.Fprintln(out, "Yes")
			} else {
				fmt.Fprintln(out, "No")
			}
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
