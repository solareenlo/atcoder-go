package main

import "fmt"

var (
	g    [9][9]bool = [9][9]bool{}
	seen [9]bool    = [9]bool{}
)

func dfs(v, n int) int {
	all := true
	for i := 0; i < n; i++ {
		if !seen[i] {
			all = false
		}
	}
	if all {
		return 1
	}

	res := 0
	for i := 0; i < n; i++ {
		if !g[v][i] || seen[i] {
			continue
		}
		seen[i] = true
		res += dfs(i, n)
		seen[i] = false
	}
	return res
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	var a, b int
	for i := 0; i < m; i++ {
		fmt.Scan(&a, &b)
		a--
		b--
		g[a][b] = true
		g[b][a] = true
	}
	seen[0] = true
	fmt.Println(dfs(0, n))
}
