package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 300003

type pair struct{ x, y int }

var (
	n int
	m int
	c = make([]int, N)
	a = make([][]pair, N)
)

func add(x, v int) {
	for i := x; i <= m; i += (i & -i) {
		c[i] += v
	}
}

func qry(x int) int {
	r := 0
	for i := x; i > 0; i -= (i & -i) {
		r += c[i]
	}
	return r
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	fmt.Fscan(in, &n, &m)

	for i := 1; i <= n; i++ {
		var l, r int
		fmt.Fscan(in, &l, &r)
		a[r-l+1] = append(a[r-l+1], pair{l, r})
	}

	cnt := n
	for i := 1; i <= m; i++ {
		r := cnt
		for o := i; o <= m; o += i {
			r += qry(o)
		}
		fmt.Fprintln(out, r)
		for _, o := range a[i] {
			add(o.x, 1)
			add(o.y+1, -1)
		}
		cnt -= len(a[i])
	}
}
