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

	var n, q int
	fmt.Fscan(in, &n, &q)

	p := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &p[i])
	}
	ps := make([]int, n)
	copy(ps, p)
	t := make([]int, q)
	a := make([]int, q)
	x := make([]int, q)
	for i := 0; i < q; i++ {
		fmt.Fscan(in, &t[i], &a[i])
		if t[i] == 1 {
			fmt.Fscan(in, &x[i])
			ps = append(ps, x[i])
		}
	}

	sort.Ints(ps)
	ps = unique(ps)
	for i := 0; i < n; i++ {
		p[i] = lowerBound(ps, p[i])
	}
	for i := 0; i < q; i++ {
		if t[i] == 1 {
			x[i] = lowerBound(ps, x[i])
		}
	}

	F := NewFenwick(len(ps))
	for i := 0; i < n; i++ {
		F.Add(p[i], 1)
	}

	ind := make([]int, len(ps))
	for i := range ind {
		ind[i] = -1
	}
	for i := 0; i < n; i++ {
		ind[p[i]] = i
	}

	for i := 0; i < q; i++ {
		if t[i] == 1 {
			a[i]--
			F.Add(p[a[i]], -1)
			ind[p[a[i]]] = -1
			p[a[i]] = x[i]
			F.Add(p[a[i]], 1)
			ind[p[a[i]]] = a[i]
		}
		if t[i] == 2 {
			a[i]--
			fmt.Fprintln(out, F.Sum(p[a[i]], len(ps)))
		}
		if t[i] == 3 {
			ok := 0
			ng := len(ps)
			for ng-ok > 1 {
				mid := (ok + ng) / 2
				if F.Sum(mid, len(ps)) >= a[i] {
					ok = mid
				} else {
					ng = mid
				}
			}
			fmt.Fprintln(out, ind[ok]+1)
		}
	}
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
	// sort.Ints(result)
	n := len(result)
	for i := 0; i < n; i++ {
		a[i] = result[i]
	}
	return result
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
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
