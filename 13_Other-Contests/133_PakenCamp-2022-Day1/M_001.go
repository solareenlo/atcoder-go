package main

import (
	"bufio"
	"fmt"
	"os"
)

var G [1 << 17][]int
var A [1 << 17]bool
var uf *dsu

func main() {
	in := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(in, &N)
	uf = NewDsu(N + 2)
	for i := 1; i < N; i++ {
		var p int
		fmt.Fscan(in, &p)
		G[p-1] = append(G[p-1], i)
	}
	var K int
	fmt.Fscan(in, &K)
	for i := 0; i < K; i++ {
		var m, s int
		fmt.Fscan(in, &m, &s)
		m--
		uf.Merge(m, N+s)
	}
	for i := 0; i < N; i++ {
		var a int
		fmt.Fscan(in, &a)
		a--
		A[a] = true
		uf.Merge(i, a)
	}
	dfs(0, -1)
	if uf.Same(N, N+1) {
		fmt.Println("IMPOSSIBLE")
	} else {
		for i := 0; i < N; i++ {
			if uf.Same(N, i) {
				fmt.Print(0)
			} else {
				fmt.Print(1)
			}
		}
		fmt.Println()
	}
}

func dfs(u, p int) {
	if p != -1 {
		uf.Merge(u, p)
	}
	if A[u] {
		p = u
	}
	for _, v := range G[u] {
		dfs(v, p)
	}
}

type dsu struct {
	n            int
	parentOrSize []int
}

func NewDsu(n int) *dsu {
	d := new(dsu)
	d.n = n
	d.parentOrSize = make([]int, d.n)
	for i := range d.parentOrSize {
		d.parentOrSize[i] = -1
	}
	return d
}

func (d *dsu) Merge(a, b int) int {
	if !(0 <= a && a < d.n) {
		panic("")
	}
	if !(0 <= b && b < d.n) {
		panic("")
	}
	x := d.Leader(a)
	y := d.Leader(b)
	if x == y {
		return x
	}
	if -d.parentOrSize[x] < -d.parentOrSize[y] {
		x, y = y, x
	}
	d.parentOrSize[x] += d.parentOrSize[y]
	d.parentOrSize[y] = x
	return x
}

func (d *dsu) Same(a, b int) bool {
	if !(0 <= a && a < d.n) {
		panic("")
	}
	if !(0 <= b && b < d.n) {
		panic("")
	}
	return d.Leader(a) == d.Leader(b)
}

func (d *dsu) Leader(a int) int {
	if !(0 <= a && a < d.n) {
		panic("")
	}
	if d.parentOrSize[a] < 0 {
		return a
	}
	d.parentOrSize[a] = d.Leader(d.parentOrSize[a])
	return d.parentOrSize[a]
}

func (d *dsu) Size(a int) int {
	if !(0 <= a && a < d.n) {
		panic("")
	}
	return -d.parentOrSize[d.Leader(a)]
}

func (d *dsu) Groups() [][]int {
	leaderBuf := make([]int, d.n)
	groupSize := make([]int, d.n)
	for i := 0; i < d.n; i++ {
		leaderBuf[i] = d.Leader(i)
		groupSize[leaderBuf[i]]++
	}
	result := make([][]int, d.n)
	for i := 0; i < d.n; i++ {
		result[i] = make([]int, 0, groupSize[i])
	}
	for i := 0; i < d.n; i++ {
		result[leaderBuf[i]] = append(result[leaderBuf[i]], i)
	}
	eraseEmpty := func(a [][]int) [][]int {
		result := make([][]int, 0, len(a))
		for i := range a {
			if len(a[i]) != 0 {
				result = append(result, a[i])
			}
		}
		return result
	}
	return eraseEmpty(result)
}
