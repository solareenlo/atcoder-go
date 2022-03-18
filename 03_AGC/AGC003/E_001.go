package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const N = 100005

var (
	m int
	a = make([]int, N)
	c = make([]int, N)
	d = make([]int, N)
)

func calc(p, x, t int) {
	d[p] += x / a[p] * t
	x %= a[p]
	p = upperBound(a[1:1+m], x)
	if p != 0 {
		calc(p, x, t)
	} else {
		c[1] += t
		c[x+1] -= t
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var i, n, q int
	fmt.Fscan(in, &n, &q)
	m = 1
	a[m] = n
	for i = 1; i <= q; i++ {
		var x int
		fmt.Fscan(in, &x)
		for m != 0 && a[m] >= x {
			m--
		}
		m++
		a[m] = x
	}

	d[m] = 1
	for i = m; i > 1; i-- {
		calc(i-1, a[i], d[i])
	}
	for i = 1; i <= a[1]; i++ {
		c[i] += c[i-1]
		fmt.Fprintln(out, c[i]+d[1])
	}
	for ; i <= n; i++ {
		fmt.Fprintln(out, 0)
	}
}

func upperBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] > x
	})
	return idx
}
