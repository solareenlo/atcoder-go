package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const M = 1000000007

	var n, k int
	var edge [][]int

	fmt.Fscan(in, &n, &k)
	edge = make([][]int, n)
	for i := 0; i < n-1; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		a--
		b--
		edge[a] = append(edge[a], b)
		edge[b] = append(edge[b], a)
	}
	var dfs func(int, int) [][]int
	dfs = func(p, par int) [][]int {
		dp := make([][]int, k+1)
		for i := range dp {
			dp[i] = make([]int, 3)
		}
		dp[0][0] = 1
		for _, c := range edge[p] {
			if c == par {
				continue
			}
			a := dfs(c, p)
			nex := make([][]int, k+1)
			for i := range nex {
				nex[i] = make([]int, 3)
			}
			for i := 0; i <= k; i++ {
				for j := 0; j < 3; j++ {
					for l := 0; l <= i; l++ {
						for m := 0; m <= j; m++ {
							nex[i][j] = (nex[i][j] + dp[l][m]*a[i-l][j-m]) % M
						}
					}
				}
			}
			dp = nex
		}
		for i := k; i >= 0; i-- {
			dp[i][2] = 0
			dp[i][1] = (dp[i][0] + dp[i][1]) % M
			if i != 0 {
				dp[i][0] = (dp[i][0] + dp[i-1][1] + dp[i-1][2]) % M
			}
		}
		return dp
	}
	fmt.Println(dfs(0, -1)[k][0])
}
