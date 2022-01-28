package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n+1)
	b := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&a[i])
	}
	for i := 1; i <= n; i++ {
		fmt.Scan(&b[i])
	}

	mod := int(1e9 + 7)
	C := make([]int, n+1)
	C[0] = 1
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			C[i] = (C[i] + C[j]*C[i-j-1]) % mod
		}
	}

	dp := [1010][1010]int{}
	dp[1][1] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if i > 1 || j > 1 {
				dp[i][j] = (dp[i-1][j] + dp[i][j-1]) % mod
			}
			if a[i] == b[j] {
				for k, cnt := 1, 0; k < i && k < j; k++ {
					if a[i-k] == b[j-k] {
						dp[i][j] = (dp[i][j] - C[cnt]*dp[i-k][j-k]%mod + mod) % mod
						cnt++
					}
				}
			}
		}
	}

	fmt.Println(dp[n][n])
}
