package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const INF = int(1e18)

	var cp [30][200][200]bool
	var dp, ep [30][200][200]int
	var n, m, W, S, K int
	fmt.Fscan(in, &n, &m, &W, &S, &K)
	S--
	for m > 0 {
		m--
		var u, v, w int
		fmt.Fscan(in, &u, &v, &w)
		u--
		v--
		cp[0][u][v] = true
		dp[0][u][v] = w
		ep[0][u][v] = max(0, w)
	}
	for h := 1; h < 30; h++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				dp[h][i][j] = -INF
			}
		}
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if !cp[h-1][i][j] {
					continue
				}
				for k := 0; k < n; k++ {
					if !cp[h-1][j][k] {
						continue
					}
					cp[h][i][k] = true
					dp[h][i][k] = max(dp[h][i][k], dp[h-1][i][j]+dp[h-1][j][k])
					ep[h][i][k] = max(ep[h][i][k], ep[h-1][i][j]+dp[h-1][j][k])
					ep[h][i][k] = max(ep[h][i][k], ep[h-1][j][k])
				}
			}
		}
	}
	var fp [200]int
	var fcp [200]bool
	fp[S] = W
	fcp[S] = true
	for h := 0; h < 30; h++ {
		if K&(1<<h) == 0 {
			continue
		}
		var gp [200]int
		var gcp [200]bool
		for i := 0; i < n; i++ {
			if !fcp[i] {
				continue
			}
			for j := 0; j < n; j++ {
				if !cp[h][i][j] {
					continue
				}
				gcp[j] = true
				gp[j] = max(gp[j], fp[i]+dp[h][i][j])
				gp[j] = max(gp[j], ep[h][i][j])
			}
		}
		for i := 0; i < n; i++ {
			fp[i] = gp[i]
			fcp[i] = gcp[i]
		}
	}
	ans := -1
	for i := 0; i < n; i++ {
		if fcp[i] {
			ans = max(ans, fp[i])
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
