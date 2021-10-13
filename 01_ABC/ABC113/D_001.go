package main

import "fmt"

func main() {
	var h, w, k int
	fmt.Scan(&h, &w, &k)
	dp := make([][]int, h+1)
	for i := range dp {
		dp[i] = make([]int, w)
	}
	dp[0][0] = 1

	mod := int(1e9 + 7)
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			for bit := 0; bit < 1<<(w-1); bit++ {
				ok := true
				for l := 0; l < w-2; l++ {
					if (bit>>l)&1 != 0 && (bit>>(l+1))&1 != 0 {
						ok = false
					}
				}
				if ok {
					if j >= 1 && (bit>>(j-1))&1 != 0 {
						dp[i+1][j-1] = (dp[i+1][j-1] + dp[i][j]) % mod
					} else if j <= w-2 && (bit>>j)&1 != 0 {
						dp[i+1][j+1] = (dp[i+1][j+1] + dp[i][j]) % mod
					} else {
						dp[i+1][j] = (dp[i+1][j] + dp[i][j]) % mod
					}
				}
			}
		}
	}
	fmt.Println(dp[h][k-1])
}
