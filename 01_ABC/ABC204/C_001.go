package main

import "fmt"

var n, m, a, b, res int

var g = make([][]int, 2001)

func dfs(u int, vis *[]bool) {
	if (*vis)[u] {
		return
	}
	(*vis)[u] = true
	res++
	for _, v := range g[u] {
		dfs(v, vis)
	}
}

func main() {
	fmt.Scan(&n, &m)
	for i := 0; i < m; i++ {
		fmt.Scan(&a, &b)
		g[a] = append(g[a], b)
	}
	for i := 1; i <= n; i++ {
		var vis = make([]bool, n+1)
		dfs(i, &vis)
	}
	fmt.Println(res)
}
