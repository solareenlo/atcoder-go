package main

import "fmt"

var (
	n, m, q int
	a       = make([]int, 0)
	b       = make([]int, 0)
	c       = make([]int, 0)
	d       = make([]int, 0)
	res     int
)

func dfs(g []int) {
	if len(g) == n+1 {
		sum := 0
		for i := 0; i < q; i++ {
			if g[b[i]]-g[a[i]] == c[i] {
				sum += d[i]
			}
		}
		res = max(res, sum)
		return
	}
	g = append(g, g[len(g)-1])
	for g[len(g)-1] <= m {
		dfs(g)
		g[len(g)-1]++
	}
}

func main() {
	fmt.Scan(&n, &m, &q)

	a = make([]int, q)
	b = make([]int, q)
	c = make([]int, q)
	d = make([]int, q)
	for i := 0; i < q; i++ {
		fmt.Scan(&a[i], &b[i], &c[i], &d[i])
	}

	dfs([]int{1})
	fmt.Println(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
