package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type unionFind struct {
	v []int
}

func newUnionFind(n int) *unionFind {
	v := make([]int, n)
	for i := range v {
		v[i] = -1
	}
	return &unionFind{v}
}

func (uf *unionFind) find(x int) int {
	if uf.v[x] < 0 {
		return x
	}
	uf.v[x] = uf.find(uf.v[x])
	return uf.v[x]
}

func (uf *unionFind) unite(x, y int) {
	x = uf.find(x)
	y = uf.find(y)
	if x != y {
		if uf.v[x] > uf.v[y] {
			x, y = y, x
		}
		uf.v[x] += uf.v[y]
		uf.v[y] = x
	}
}

func (uf *unionFind) same(x, y int) bool {
	return uf.find(x) == uf.find(y)
}

func (uf *unionFind) size(x int) int {
	return -uf.v[uf.find(x)]
}

func f(n int, a, b, p, c, d, q []int) int64 {
	uf := newUnionFind(n)
	_uf := newUnionFind(n)
	num := make([]map[int]int, n)
	for u := 0; u < n; u++ {
		num[u] = make(map[int]int)
		num[u][u] = 1
	}
	C := make([][]int, n)
	for u := 0; u < n; u++ {
		C[u] = []int{u}
	}
	j := 0
	var sum int64
	for i := 0; i < n-1; i++ {
		for ; j < n-1 && q[j] <= p[i]; j++ {
			c[j] = _uf.find(c[j])
			d[j] = _uf.find(d[j])
			if _uf.size(c[j]) < _uf.size(d[j]) {
				c[j], d[j] = d[j], c[j]
			}
			_uf.unite(c[j], d[j])
			for _, k := range C[d[j]] {
				C[c[j]] = append(C[c[j]], k)
			}
			vis := make(map[int]bool)
			for _, k := range C[d[j]] {
				x := uf.find(k)
				if vis[x] {
					continue
				}
				vis[x] = true
				num[x][c[j]] += num[x][d[j]]
				delete(num[x], d[j])
			}
			C[d[j]] = nil
		}
		a[i] = uf.find(a[i])
		b[i] = uf.find(b[i])
		if uf.size(a[i]) < uf.size(b[i]) {
			a[i], b[i] = b[i], a[i]
		}
		tmp := int64(uf.size(a[i]) * uf.size(b[i]))
		uf.unite(a[i], b[i])
		for x, y := range num[b[i]] {
			tmp -= int64(num[a[i]][x] * y)
			num[a[i]][x] += y
		}
		num[b[i]] = nil
		sum += tmp
	}
	return sum
}

type pair struct {
	first, second int
}

type P struct {
	first  int
	second pair
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(in, &N)
	a := make([]int, N-1)
	b := make([]int, N-1)
	p := make([]int, N-1)
	for i := 0; i < N-1; i++ {
		fmt.Fscan(in, &a[i], &b[i], &p[i])
		a[i]--
		b[i]--
	}
	c := make([]int, N-1)
	d := make([]int, N-1)
	q := make([]int, N-1)
	for i := 0; i < N-1; i++ {
		fmt.Fscan(in, &c[i], &d[i], &q[i])
		c[i]--
		d[i]--
	}
	pab := make([]P, N-1)
	qcd := make([]P, N-1)
	for i := 0; i < N-1; i++ {
		pab[i] = P{p[i], pair{a[i], b[i]}}
		qcd[i] = P{q[i], pair{c[i], d[i]}}
	}
	sort.Slice(pab, func(i, j int) bool {
		return pab[i].first < pab[j].first
	})
	sort.Slice(qcd, func(i, j int) bool {
		return qcd[i].first < qcd[j].first
	})
	for i := 0; i < N-1; i++ {
		a[i] = pab[i].second.first
		b[i] = pab[i].second.second
		p[i] = pab[i].first
		c[i] = qcd[i].second.first
		d[i] = qcd[i].second.second
		q[i] = qcd[i].first
	}
	ans := int64(N * (N - 1) / 2)
	ans -= f(N, a, b, p, c, d, q)
	ans -= f(N, c, d, q, a, b, p)
	fmt.Println(ans)
}
