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

	var n int
	fmt.Fscan(in, &n)

	g := make([][][2]int, n+1)
	for i := 1; i < n; i++ {
		var a, b, c int
		fmt.Fscan(in, &a, &b, &c)
		g[a] = append(g[a], [2]int{b, c})
		g[b] = append(g[b], [2]int{a, c})
	}

	d := make([]int, n)
	l := make([]int, n)
	for i := range d {
		fmt.Fscan(in, &d[i])
	}

	var dfs func(u, v, x int) [2]int
	dfs = func(u, v, x int) [2]int {
		s := v
		e := d[v-1]
		l[v-1] = max(l[v-1], x)
		if x == 0 {
			x = e
			e = 0
		}
		for i := range g[v] {
			w := g[v][i][0]
			c := g[v][i][1]
			if w != u {
				res := dfs(v, w, x+c)
				S := res[0]
				D := res[1]
				if e < c+D {
					e = c + D
					s = S
				}
			}
		}
		return [2]int{s, e}
	}

	dfs(0, dfs(0, dfs(0, 1, 0)[0], 0)[0], 0)
	for i := range l {
		fmt.Fprintln(out, l[i])
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
