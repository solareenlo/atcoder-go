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
	a := make([]int, k+1)
	for i := 1; i <= k; i++ {
		fmt.Fscan(in, &a[i])
	}

	const mod = 1_000_000_007
	c := [1001][1001]int{}
	c[0][0] = 1
	for i := 1; i <= n; i++ {
		c[i][0] = 1
		for j := 1; j <= i; j++ {
			c[i][j] = (c[i-1][j] + c[i-1][j-1]) % mod
		}
	}

	dp := [21][1001]int{}
	dp[0][0] = 1
	for i := 1; i <= k; i++ {
		for j := 0; j <= n; j++ {
			if dp[i-1][j] != 0 {
				for k := 0; j+k <= n && k <= a[i]; k++ {
					dp[i][j+k] = (dp[i][j+k] + dp[i-1][j]*c[j+k][j]%mod*c[n-k][a[i]-k]%mod) % mod
				}
			}
		}
	}
	fmt.Println(dp[k][n])
}
