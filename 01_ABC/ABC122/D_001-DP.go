package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, 64)
	}
	dp[0][63] = 1

	mod := int(1e9 + 7)
	for i := 0; i < n; i++ {
		for j := 0; j < 64; j++ {
			a, b, c := j/16, j/4%4, j%4
			for d := 0; d < 4; d++ {
				if b == 0 && c == 2 && d == 1 {
					continue
				}
				if a == 0 && c == 2 && d == 1 {
					continue
				}
				if b == 2 && c == 0 && d == 1 {
					continue
				}
				if b == 0 && c == 1 && d == 2 {
					continue
				}
				if a == 0 && b == 2 && d == 1 {
					continue
				}
				dp[i+1][b*16+c*4+d] += dp[i][j]
				dp[i+1][b*16+c*4+d] %= mod
			}
		}
	}

	sum := 0
	for i := 0; i < 64; i++ {
		sum += dp[n][i]
		sum %= mod
	}
	fmt.Println(sum)
}
