package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	g := make([][]int, n)
	for i := 0; i < n-1; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		a--
		b--
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}

	v := make([]bool, n)
	dp := make([]int, n)
	var dfs func(c int)
	dfs = func(c int) {
		v[c] = true
		dp[c] = 1
		for _, x := range g[c] {
			if v[x] {
				continue
			}
			dfs(x)
			dp[c] += dp[x]
		}
	}

	dfs(0)

	ans := 0
	for i := 0; i < n; i++ {
		ans += dp[i] * (n - dp[i])
	}
	fmt.Println(ans)
}
