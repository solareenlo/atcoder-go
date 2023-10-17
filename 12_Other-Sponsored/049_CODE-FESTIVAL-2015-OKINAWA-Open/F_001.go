package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const maxN = 22
const inf = 1e9

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	var x, y [maxN]float64
	var d [maxN][maxN]float64
	var dp [maxN][1 << maxN]float64

	// Read input
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x[i], &y[i])
	}

	// Initialize dp
	for i := 0; i <= n; i++ {
		for j := 0; j < 1<<n; j++ {
			dp[i][j] = inf
		}
	}
	dp[n][0] = 0.0

	for i := 0; i < 3; i++ {
		fmt.Fscan(in, &x[n], &y[n])

		// Copy previous dp values
		for j := 0; j < 1<<n; j++ {
			cur := inf
			for k := 0; k <= n; k++ {
				cur = math.Min(cur, dp[k][j])
				dp[k][j] = inf
			}
			dp[n][j] = cur
		}

		// Calculate distances
		for j := 0; j <= n; j++ {
			for k := 0; k <= n; k++ {
				d[j][k] = math.Hypot(x[j]-x[k], y[j]-y[k])
			}
		}

		// Update dp using dynamic programming
		for j := 0; j < 1<<n; j++ {
			for k := 0; k < n; k++ {
				if (j>>k)&1 != 0 {
					sub := j - (1 << k)
					for l := 0; l <= n; l++ {
						if l == n || (sub>>l)&1 != 0 {
							dp[k][j] = math.Min(dp[k][j], dp[l][sub]+d[k][l])
						}
					}
				}
			}
		}
	}

	// Find the minimum path
	ret := inf
	for i := 0; i <= n; i++ {
		ret = math.Min(ret, dp[i][(1<<n)-1])
	}

	fmt.Printf("%.15f\n", ret)
}
