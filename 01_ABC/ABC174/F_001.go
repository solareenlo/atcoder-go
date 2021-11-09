package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	n, q int
	b    = [500005]int{}
)

func add(x, v int) {
	for i := x; i <= n; i += i & (-i) {
		b[i] += v
	}
}

func sum(x int) int {
	v := 0
	for i := x; i > 0; i -= i & (-i) {
		v += b[i]
	}
	return v
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	fmt.Fscan(in, &n, &q)

	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}

	l := make([]int, q+1)
	r := make([]int, q+1)
	type pair struct{ u, v int }
	x := make([][]pair, 500005)
	for i := 1; i <= q; i++ {
		fmt.Fscan(in, &l[i], &r[i])
		x[r[i]] = append(x[r[i]], pair{l[i], i})
	}

	c := make([]int, 500005)
	res := make([]int, 500005)
	for i := 1; i <= n; i++ {
		if c[a[i]] != 0 {
			add(c[a[i]], -1)
		}
		c[a[i]] = i
		add(i, 1)
		for _, p := range x[i] {
			res[p.v] = sum(i) - sum(p.u-1)
		}
	}

	for i := 1; i <= q; i++ {
		fmt.Fprintln(out, res[i])
	}
}
