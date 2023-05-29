package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var idx, val [2 << 17]int
	var G [2 << 17][]int

	var N, P, Q int
	fmt.Fscan(in, &N, &P, &Q)
	uf := NewDsu(N)
	for i := 0; i < P; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		a--
		b--
		uf.Merge(a, b)
	}
	for i := 0; i < N; i++ {
		idx[i] = -1
	}
	sz := 0
	for i := 0; i < N; i++ {
		if idx[uf.Leader(i)] == -1 {
			idx[uf.Leader(i)] = sz
			sz++
		}
	}
	for i := 0; i < Q; i++ {
		var c, d int
		fmt.Fscan(in, &c, &d)
		c--
		d--
		c = idx[uf.Leader(c)]
		d = idx[uf.Leader(d)]
		if c == d {
			fmt.Println(-1)
			return
		}
		if c > d {
			c, d = d, c
		}
		G[d] = append(G[d], c)
	}
	ban := make([]int, sz)
	for i := range ban {
		ban[i] = -1
	}
	for i := 0; i < sz; i++ {
		for _, u := range G[i] {
			ban[val[u]] = i
		}
		k := 0
		for ban[k] == i {
			k++
		}
		val[i] = k
	}
	for i := 0; i < N; i++ {
		if i+1 == N {
			fmt.Println(1 + val[idx[uf.Leader(i)]])
		} else {
			fmt.Printf("%d ", 1+val[idx[uf.Leader(i)]])
		}
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
