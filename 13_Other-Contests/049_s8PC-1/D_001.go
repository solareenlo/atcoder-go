package main

import "fmt"

const MOD = 1e9 + 7

var dp [1001][1001]int

func main() {
	var h, w, k int
	fmt.Scan(&h, &w, &k)
	for i := 0; i < 1001; i++ {
		for j := 0; j < 1001; j++ {
			dp[i][j] = 1
		}
	}
	for ; k > 0; k-- {
		var x, y int
		fmt.Scan(&x, &y)
		if x != 1 {
			for i := y + 1; i < h; i++ {
				dp[i][x-2] = 0
			}
		}
		if y != 1 {
			for i := x + 1; i < w; i++ {
				dp[y-2][i] = 0
			}
		}
	}
	for i := 1; i < h; i++ {
		for j := 1; j < w; j++ {
			if dp[i][j] != 0 {
				dp[i][j] = (dp[i-1][j] + dp[i][j-1]) % MOD
			}
		}
	}
	fmt.Println(dp[h-1][w-1])
}
