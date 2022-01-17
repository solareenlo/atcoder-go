package main

import "fmt"

var (
	seen = [101]bool{}
	G    = make([][]int, 101)
	h    = [101]int{}
)

func dfs(v int) int {
	if seen[v] {
		return 0
	}
	seen[v] = true
	res := 0
	for _, nv := range G[v] {
		res += dfs(nv)
	}
	if res == 0 {
		if h[v] != 0 {
			return 2
		} else {
			return 0
		}
	}
	return res + 2
}

func main() {
	var n, x int
	fmt.Scan(&n, &x)
	x--

	for i := 0; i < n; i++ {
		fmt.Scan(&h[i])
	}
	for i := 0; i < n-1; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		a--
		b--
		G[a] = append(G[a], b)
		G[b] = append(G[b], a)
	}

	fmt.Println(max(dfs(x)-2, 0))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
