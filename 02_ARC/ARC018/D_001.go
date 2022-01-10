package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func det(a [][]int) int {
	n := len(a)
	res := 1
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if a[j][i] != 0 {
				if j != i {
					a[i], a[j] = a[j], a[i]
					res = (-res + mod) % mod
				}
				break
			}
		}
		if a[i][i] == 0 {
			return 0
		}
		res = res * a[i][i] % mod
		inv := invMod(a[i][i])
		for j := i + 1; j < n; j++ {
			k := a[j][i] * inv % mod
			for l := i; l < n; l++ {
				a[j][l] = (a[j][l] - a[i][l]*k%mod + mod) % mod
			}
		}
	}
	return res
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)
	type pair struct{ x, y int }
	link := map[int][]pair{}
	for i := 0; i < m; i++ {
		var a, b, c int
		fmt.Fscan(in, &a, &b, &c)
		link[c] = append(link[c], pair{a - 1, b - 1})
	}

	keys := make([]int, 0, len(link))
	for k := range link {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	uf := New(n)
	sum := 0
	res := 1
	for _, cost := range keys {
		l := link[cost]
		usel := make([]pair, 0)
		for _, x := range l {
			u := uf.Root(x.x)
			v := uf.Root(x.y)
			if u == v {
				continue
			}
			usel = append(usel, pair{u, v})
		}
		dl := map[int][]pair{}
		for _, x := range usel {
			u := uf.Root(x.x)
			v := uf.Root(x.y)
			if u != v {
				sum += cost
				uf.Merge(u, v)
				if _, ok := dl[v]; ok {
					dl[u] = append(dl[u], dl[v]...)
					delete(dl, v)
				}
			}
			dl[u] = append(dl[u], pair{x.x, x.y})
		}
		for _, j := range dl {
			nidx := make([]int, 0)
			for _, x := range j {
				nidx = append(nidx, x.x)
				nidx = append(nidx, x.y)
			}
			sort.Ints(nidx)
			nidx = unique(nidx)
			N := len(nidx)
			if N == 1 {
				continue
			}
			a := make([][]int, N-1)
			for i := 0; i < N-1; i++ {
				a[i] = make([]int, N-1)
			}
			for _, x := range j {
				u := lowerBound(nidx, x.x)
				v := lowerBound(nidx, x.y)
				if u != N-1 && v != N-1 {
					a[u][v] -= 1
					a[v][u] -= 1
				}
				if u != N-1 {
					a[u][u] += 1
				}
				if v != N-1 {
					a[v][v] += 1
				}
			}
			for i := 0; i < N-1; i++ {
				for j := 0; j < N-1; j++ {
					a[i][j] = (a[i][j] + mod) % mod
				}
			}
			res = (res * det(a)) % mod
		}
	}
	fmt.Println(sum, res)
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
	return result
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}

const mod = 1000000007

func powMod(a, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = res * a % mod
		}
		a = a * a % mod
		n /= 2
	}
	return res
}

func invMod(a int) int {
	return powMod(a, mod-2)
}

type UnionFind struct {
	root []int
	n    int
}

func New(size int) *UnionFind {
	uf := new(UnionFind)
	uf.root = make([]int, size)
	uf.n = size
	for i := 0; i < size; i++ {
		uf.root[i] = i
	}
	return uf
}

func (uf *UnionFind) Merge(p, q int) bool {
	q = uf.Root(q)
	p = uf.Root(p)
	if q == p {
		return false
	}
	uf.root[q] = p
	return true
}

func (uf *UnionFind) Root(p int) int {
	if uf.root[p] == p {
		return p
	}
	uf.root[p] = uf.Root(uf.root[p])
	return uf.root[p]
}
