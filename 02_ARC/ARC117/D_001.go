package main

import (
	"bufio"
	"fmt"
	"os"
)

const M = 1 << 18

var (
	F = [M]int{}
	p = [M]int{}
	E = [M]int{}
	d = [M]int{}
	D int
	G = make([][]int, M)
)

func dfs(x, f int, s *int) {
	F[x] = f
	d[x] = d[f] + 1
	if d[x] > d[*s] {
		*s = x
	}
	for _, y := range G[x] {
		if y^f != 0 {
			dfs(y, x, s)
		}
	}
}

func dfs2(u, f int) {
	D++
	E[u] = D
	for _, v := range G[u] {
		if v != f && p[v] == 0 {
			dfs2(v, u)
		}
	}
	for _, v := range G[u] {
		if v != f && p[v] != 0 {
			dfs2(v, u)
		}
	}
	D++
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	for i := 1; i < n; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		G[u] = append(G[u], v)
		G[v] = append(G[v], u)
	}

	s := 0
	dfs(1, 0, &s)
	t := 0
	dfs(s, 0, &t)

	for u := t; u > 0; u = F[u] {
		p[u] = 1
	}

	dfs2(s, 0)

	for i := 1; i <= n; i++ {
		fmt.Fprint(out, E[i], " ")
	}
	fmt.Fprintln(out)
}
