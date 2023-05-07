package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)

	g := make([][]int, n+m)
	d := make([]int, n+m)
	for i := range d {
		d[i] = -2
	}
	for m > 0 {
		m--
		var k int
		fmt.Fscan(in, &k)
		for k > 0 {
			k--
			var r int
			fmt.Fscan(in, &r)
			g[r-1] = append(g[r-1], n+m)
			g[n+m] = append(g[n+m], r-1)
		}
	}

	q := make([]int, 0)
	d[0] = 0
	q = append(q, d[0])
	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		for _, v := range g[u] {
			if d[v] < 0 {
				d[v] = d[u] + 1
				q = append(q, v)
			}
		}
	}

	for i := 0; i < n; i++ {
		fmt.Fprintln(out, d[i]/2)
	}
}
