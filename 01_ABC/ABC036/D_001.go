package main

import "fmt"

var (
	g   [][]int
	dpB []int
	dpW []int
	mod int = int(1e9 + 7)
)

func dfs(pos, q int) {
	dpB[pos] = 1
	dpW[pos] = 1
	for _, p := range g[pos] {
		if p != q {
			dfs(p, pos)
			dpB[pos] = dpB[pos] * dpW[p] % mod
			dpW[pos] = dpW[pos] * (dpB[p] + dpW[p]) % mod
		}
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	g = make([][]int, n)
	for i := 0; i < n-1; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		a--
		b--
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}

	dpB = make([]int, n)
	dpW = make([]int, n)
	dfs(0, 0)
	fmt.Println((dpB[0] + dpW[0]) % mod)
}
