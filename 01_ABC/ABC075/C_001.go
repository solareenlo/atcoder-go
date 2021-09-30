package main

import "fmt"

var (
	g    [][]int
	seen []bool
)

func dfs(v int) {
	seen[v] = true
	for _, u := range g[v] {
		if seen[u] {
			continue
		}
		dfs(u)
	}
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	a, b := make([]int, m), make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&a[i], &b[i])
		a[i]--
		b[i]--
	}

	res := 0
	for i := 0; i < m; i++ {
		g = make([][]int, n)
		for j := 0; j < m; j++ {
			if i != j {
				g[a[j]] = append(g[a[j]], b[j])
				g[b[j]] = append(g[b[j]], a[j])
			}
		}

		seen = make([]bool, n)
		dfs(a[i])
		if !seen[b[i]] {
			res++
		}
	}
	fmt.Println(res)
}
