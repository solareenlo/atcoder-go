package main

import (
	"bufio"
	"fmt"
	"os"
)

type edge struct{ v, c int }

var (
	s   int
	n   int
	mn  int
	siz = make([]int, 100005)
	g   = make([][]edge, 100005)
)

func Dfs(u, fa int) {
	p := 0
	siz[u] = 1
	for i := 0; i < len(g[u]); i++ {
		v := g[u][i].v
		if v != fa {
			Dfs(v, u)
			siz[u] += siz[v]
			if siz[v] > siz[p] {
				p = v
			}
			s += 2 * g[u][i].c * min(siz[v], n-siz[v])
		}
	}
	x := max(siz[p], n-siz[u])
	if n-siz[u] > siz[p] {
		p = fa
	}
	if x*2 <= n {
		for i := 0; i < len(g[u]); i++ {
			if x*2 != n || g[u][i].v == p {
				mn = min(mn, g[u][i].c)
			}
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n)

	for i := 1; i < n; i++ {
		var u, v, c int
		fmt.Fscan(in, &u, &v, &c)
		g[u] = append(g[u], edge{v, c})
		g[v] = append(g[v], edge{u, c})
	}

	mn = 1 << 60
	Dfs(1, 0)
	fmt.Println(s - mn)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
