package main

import "fmt"

var (
	n   int
	tot int
	cnt int
	a   = [10005][10005]bool{}
	b   = [10005]int{}
	c   = [10005]int{}
	vis = [10005]bool{}
)

func dfs1(u int) {
	if vis[u] {
		return
	}
	vis[u] = true
	b[tot] = u
	tot++
	for i := 1; i <= n; i++ {
		if a[u][i] {
			dfs1(i)
		}
	}
}

func dfs2(u int) {
	if u == tot {
		cnt++
		return
	}
	g := [3]int{}
	for i := 0; i < u; i++ {
		if a[b[u]][b[i]] {
			g[c[i]] = 1
		}
	}
	for i := 0; i < 3; i++ {
		if g[i] != 0 {
			continue
		}
		c[u] = i
		dfs2(u + 1)
	}
}

func main() {
	var m int
	fmt.Scan(&n, &m)

	for i := 0; i < m; i++ {
		var u, v int
		fmt.Scan(&u, &v)
		a[u][v] = true
		a[v][u] = true
	}

	res := 1
	for i := 1; i <= n; i++ {
		if vis[i] {
			continue
		}
		cnt, tot = 0, 0
		dfs1(i)
		dfs2(0)
		res *= cnt
	}
	fmt.Println(res)
}
