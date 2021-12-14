package main

import "fmt"

const mod = 998244353

type pair struct{ to, i int }

var (
	g = make([][]pair, 1001)
	c = make([]int, 1001)
)

func dfs(v, pre, goal int) bool {
	if v == goal {
		return true
	}
	for _, e := range g[v] {
		if e.to != pre {
			if dfs(e.to, v, goal) {
				c[e.i]++
				return true
			}
		}
	}
	return false
}

func main() {
	var n, m, k int
	fmt.Scan(&n, &m, &k)

	a := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&a[i])
		a[i]--
	}

	for i := 0; i < n-1; i++ {
		var u, v int
		fmt.Scan(&u, &v)
		u--
		v--
		g[u] = append(g[u], pair{v, i})
		g[v] = append(g[v], pair{u, i})
	}

	for i := 0; i < m-1; i++ {
		dfs(a[i], -1, a[i+1])
	}

	sum := 0
	for i := 0; i < n-1; i++ {
		sum += c[i]
	}
	if (sum+k)%2 != 0 || sum+k < 0 {
		fmt.Println(0)
		return
	}

	mod := 998244353
	dp := make([]int, 100001)
	dp[0] = 1
	for i := 0; i < n-1; i++ {
		for x := 100000; x >= c[i]; x-- {
			dp[x] += dp[x-c[i]]
			dp[x] %= mod
		}
	}
	fmt.Println(dp[(sum+k)/2])
}
