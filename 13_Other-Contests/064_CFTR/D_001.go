package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Scan(&n, &m)

	dsu := New(n)
	var a, b int
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &a, &b)
		dsu.Merge(a-1, b-1)
	}
	a, b = dsu.Size(0), dsu.Size(1)
	if a > b {
		a, b = b, a
	}
	b = n - a
	fmt.Println(a*(a-1)/2 + b*(b-1)/2 - m)
}

type DisjointSetUnion struct {
	n            int
	parentOrSize []int
}

func New(n int) *DisjointSetUnion {
	d := &DisjointSetUnion{
		n:            n,
		parentOrSize: make([]int, n),
	}
	for i := range d.parentOrSize {
		d.parentOrSize[i] = -1
	}
	return d
}

func (d *DisjointSetUnion) Merge(a, b int) int {
	if !(0 <= a && a < d.n) {
		panic("")
	}
	if !(0 <= b && b < d.n) {
		panic("")
	}
	x, y := d.Leader(a), d.Leader(b)
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

func (d *DisjointSetUnion) Leader(a int) int {
	if !(0 <= a && a < d.n) {
		panic("")
	}
	if d.parentOrSize[a] < 0 {
		return a
	}
	d.parentOrSize[a] = d.Leader(d.parentOrSize[a])
	return d.parentOrSize[a]
}

func (d *DisjointSetUnion) Size(a int) int {
	if !(0 <= a && a < d.n) {
		panic("")
	}
	return -d.parentOrSize[d.Leader(a)]
}
