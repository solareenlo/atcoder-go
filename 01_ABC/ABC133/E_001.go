package main

import "fmt"

var (
	n, k int
	g        = make([][]int, 100005)
	res  int = 1
	mod  int = int(1e9 + 7)
)

func dfs(now, from, d int) {
	res = res * max(k-d, 0) % mod
	cnt := 1
	if from != 0 {
		cnt += 1
	}
	for _, e := range g[now] {
		if e != from {
			dfs(e, now, cnt)
			cnt++
		}
	}
}

func main() {
	fmt.Scan(&n, &k)
	for i := 0; i < n-1; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}
	dfs(1, 0, 0)
	fmt.Println(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
