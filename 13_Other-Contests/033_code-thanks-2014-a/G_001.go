package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscan(in, &n, &k)

	p := make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &p[i])
	}

	dp := make([][][]float64, n+2)
	for i := range dp {
		dp[i] = make([][]float64, k+2)
		for j := range dp[i] {
			dp[i][j] = make([]float64, k+2)
		}
	}
	dp[0][0][0] = 1.0
	for i := 1; i < n; i++ {
		for j := 0; j < k; j++ {
			if j+1 < k {
				dp[i][j+1][0] += dp[i-1][j][0] * p[i] / 100.0
			} else {
				dp[i][j][0] += dp[i-1][j][0] * p[i] / 100.0
			}
			if j+2 < k {
				dp[i][j+2][1] += dp[i-1][j][0] * (100 - p[i]) / 100.0
			} else {
				dp[i][j][0] += dp[i-1][j][0] * (100 - p[i]) / 100.0
			}
			for l := 1; l < k; l++ {
				dp[i][j][l-1] += dp[i-1][j][l] * p[i] / 100.0
				if j+2 < k {
					dp[i][j+2][l+1] += dp[i-1][j][l] * (100 - p[i]) / 100.0
				} else {
					dp[i][j][l] += dp[i-1][j][l] * (100 - p[i]) / 100.0
				}
			}
		}
	}

	ans := 0.0
	for j := 0; j < k; j++ {
		for l := 0; l < k; l++ {
			ans += float64(l+(k-1-j)) * dp[n-1][j][l]
		}
	}
	fmt.Println(ans)
}
