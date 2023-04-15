package main

import (
	"bufio"
	"fmt"
	"os"
)

var g [5050][]int
var h [5050]int

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	for i := 1; i < n; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		g[u] = append(g[u], i)
		h[i] = v
	}
	dp := make([]int, 0)
	dfs(0, &dp)
	fmt.Println(len(dp))
}

func dfs(cur int, dp *[]int) {
	*dp = make([]int, 1)
	for _, child := range g[cur] {
		dp2 := make([]int, 0)
		dfs(child, &dp2)
		dp_next := make([]int, len(*dp)+len(dp2)-1)
		for i := range dp_next {
			dp_next[i] = 1 << 60
		}
		for i := 0; i < len(*dp); i++ {
			for j := 0; j < len(dp2); j++ {
				dp_next[i+j] = min(dp_next[i+j], (*dp)[i]+dp2[j])
			}
		}
		*dp = dp_next
	}
	if cur != 0 {
		for i := 0; i < len(*dp); i++ {
			if (*dp)[i] > h[cur] {
				resize(&*dp, i)
				break
			}
		}
		*dp = append(*dp, h[cur])
	}
}

func resize(a *[]int, n int) {
	if len(*a) > n {
		*a = (*a)[:n]
	} else {
		n = n - len(*a)
		for i := 0; i < n; i++ {
			*a = append(*a, 0)
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
