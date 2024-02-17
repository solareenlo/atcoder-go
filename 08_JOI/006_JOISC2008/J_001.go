package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const INF = 10000

	type tuple struct {
		r    float64
		a, b int
	}

	var n int
	fmt.Fscan(in, &n)
	x := make([]float64, n)
	y := make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x[i], &y[i])
	}
	g := make([]tuple, 0)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i != j {
				g = append(g, tuple{math.Atan2(y[j]-y[i], x[j]-x[i]), i, j})
			}
		}
	}
	sort.Slice(g, func(i, j int) bool {
		return g[i].r < g[j].r
	})
	var dp [130][130]int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i != j {
				dp[i][j] = -INF
			}
		}
	}
	for i := 0; i < len(g); i++ {
		for j := 0; j < n; j++ {
			dp[j][g[i].b] = max(dp[j][g[i].b], dp[j][g[i].a]+1)
		}
	}
	ans := 0
	for i := 0; i < n; i++ {
		ans = max(ans, dp[i][i])
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
