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
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		a--
		b--
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}

	l := 0
	dp := make([]int, n)
	for i := range dp {
		dp[i] = -1
	}
	var dfs func(int, int)
	dfs = func(node, prev int) {
		if dp[node] != -1 {
			l = dp[node] - dp[prev] + 1
			return
		}
		dp[node] = dp[prev] + 1
		for _, next := range g[node] {
			if next != prev {
				dfs(next, node)
			}
		}
	}
	dfs(0, 0)
	tmp := n
	if l%2 != 0 {
		tmp = n - 1
	}
	if l == n {
		fmt.Println(2, tmp)
	} else {
		fmt.Println(1, tmp)
	}
}
