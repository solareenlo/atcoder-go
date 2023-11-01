package main

import "fmt"

func main() {
	var n, k, p int
	fmt.Scan(&n, &k, &p)

	var dp [100][100][100]int

	for j := 0; j < n; j++ {
		for h := 0; h < n; h++ {
			dp[0][j][h] = 1
		}
	}
	for i := 1; i < n; i++ {
		m := n - i
		for l := i - 1; l >= 0; l-- {
			for j := 0; j <= m; j++ {
				for h := 1; h <= m; h++ {
					dp[l][j][h] = (dp[l][j][h] + dp[l][j][h-1]) % p
				}
			}
			for j := 1; j <= m; j++ {
				for h := 0; h <= m; h++ {
					dp[l][j][h] = (dp[l][j][h] + dp[l][j-1][h]) % p
				}
			}
			for j := 0; j < m; j++ {
				for h := 0; h < m; h++ {
					dp[l][j][h] = ((dp[l][j][m]+dp[l][m][h])%p - dp[l][j][h]*2%p + p) % p
					dp[l+1][j][h] = (dp[l+1][j][h] + (dp[l][m][m]-dp[l][j][h]+p)%p) % p
				}
			}
		}
	}
	fmt.Println(dp[k][0][0])
}
