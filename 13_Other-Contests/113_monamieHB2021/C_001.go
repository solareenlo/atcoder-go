package main

import "fmt"

func main() {
	const MOD = 998244353

	var N, K int
	fmt.Scan(&N, &K)

	var dp [205][205]int
	dp[0][0] = 1

	for t := 0; t < N; t++ {
		for x := min(t*2, K); x >= 0; x-- {
			carry_ := 0
			for y := 0; y < min(t, K)+2; y++ {
				if dp[x][y] != 0 || carry_ != 0 {
					carry := dp[x][y] % MOD
					dp[x+1][y] += carry * (x + 1)
					if y != 0 {
						dp[x+2][y-1] += carry * (x + 1) * y
						dp[x+3][y-1] += carry * y
					}
					if y >= 2 {
						dp[x+4][y-2] += carry * (y * (y - 1) / 2)
					}
					dp[x][y] = carry_ + carry*(x*(x+1)/2+y)
					carry_ = carry
				}
			}
		}
	}

	ret := 0
	for j := 0; j < K/2+1; j++ {
		ret += dp[K-j*2][j]
	}
	fmt.Println(ret % MOD)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
