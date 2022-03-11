package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	p = make([]int, 100005)
	w = make([]int, 100005)
	d = make([]int, 100005)
)

func find(x int) int {
	if p[x] < 0 {
		return x
	}
	p[x] = find(p[x])
	return p[x]
}

func uni(x, y, z int) {
	x = find(x)
	y = find(y)
	if x != y {
		if p[x] > p[y] {
			x, y = y, x
		}
		p[x] += p[y]
		w[x] += w[y]
		d[x] += d[y]
		p[y] = x
	}
	if z <= w[x] {
		d[x] = 0
	} else {
		d[x]++
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &w[i])
		p[i] = -1
	}

	type T struct{ a, b, c int }
	tp := make([]T, m)
	for i := 0; i < m; i++ {
		var a, b, c int
		fmt.Fscan(in, &a, &b, &c)
		tp[i] = T{c, a - 1, b - 1}
	}
	sort.Slice(tp, func(i, j int) bool {
		return tp[i].a < tp[j].a
	})

	var a, b, c int
	for i := 0; i < m; i++ {
		c = tp[i].a
		a = tp[i].b
		b = tp[i].c
		uni(a, b, c)
	}
	fmt.Println(d[find(a)])
}
