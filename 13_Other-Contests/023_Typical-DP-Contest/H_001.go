package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const INF = int(1e9 + 2)

	type pair struct {
		x, y int
	}

	var N, W, C int
	fmt.Fscan(in, &N, &W, &C)
	dp := make([][]int, C+1)
	for i := range dp {
		dp[i] = make([]int, W+1)
		for j := range dp[i] {
			dp[i][j] = -INF
		}
	}
	dp[0][0] = 0
	cl := make([][]pair, 51)
	for i := 0; i < N; i++ {
		var w, v, c int
		fmt.Fscan(in, &w, &v, &c)
		cl[c] = append(cl[c], pair{w, v})
	}
	ndp := make([][]int, len(dp))
	for i := range ndp {
		ndp[i] = make([]int, len(dp[i]))
	}

	COPY := func() {
		for i := range ndp {
			copy(ndp[i], dp[i])
		}
	}

	for i := 0; i < 51; i++ {
		COPY()
		for _, S := range cl[i] {
			w := S.x
			v := S.y
			for k := 0; k < C; k++ {
				for j := W; j >= 0; j-- {
					if j+w <= W {
						ndp[k][j+w] = max(ndp[k][j+w], ndp[k][j]+v)
					}
				}
			}
		}
		for i := 0; i < C; i++ {
			for j := 0; j <= W; j++ {
				dp[i+1][j] = max(dp[i+1][j], ndp[i][j])
			}
		}
	}
	ans := 0
	for i := 0; i <= C; i++ {
		for j := 0; j <= W; j++ {
			ans = max(ans, dp[i][j])
		}
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
