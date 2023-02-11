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

	g := make([][]int, n)
	ans := make([]int, n)
	const inf = -1
	for i := range ans {
		ans[i] = inf
	}

	for i := 0; i < m; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		a--
		b--
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}

	ans[0] = 0
	q := make([]int, 0)
	q = append(q, 0)
	for len(q) > 0 {
		a := q[0]
		q = q[1:]
		for _, b := range g[a] {
			if ans[b] != inf {
				continue
			}
			ans[b] = ans[a] + 1
			q = append(q, b)
		}
	}

	for i := 0; i < n; i++ {
		fmt.Fprintln(out, ans[i])
	}
}
