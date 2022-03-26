package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	const mod = 1_000_000_007
	c := [55][55]int{}
	dp := [55][55][55]int{}
	dp[1][1][1] = 1
	for i := 0; i < n+1; i++ {
		c[i][0] = 1
		c[i][i] = 1
		for j := 1; j < i; j++ {
			c[i][j] = c[i-1][j] + c[i-1][j-1]
			c[i][j] %= mod
		}
	}

	for i := 1; i < n; i++ {
		for j := 1; j < i+1; j++ {
			for k := 0; k < j+1; k++ {
				for l := 0; l < j+1; l++ {
					tmp := 0
					if k-l > 0 {
						tmp = k - l
					}
					for w := tmp; w < ((n-i+k-l)>>1)+1; w++ {
						if w+l-k+w > 0 {
							dp[w+l-k+w+i][w+l-k+w][w] += dp[i][j][k] * c[j][l] % mod * c[w+l-k+w+i][i]
							dp[w+l-k+w+i][w+l-k+w][w] %= mod
						}
					}
				}
			}
		}
	}

	ans := 0
	for i := 1; i < n+1; i++ {
		for j := 0; j < i+1; j++ {
			ans += dp[n][i][j] * c[i][j]
			ans %= mod
		}
	}
	fmt.Println(ans)
}
