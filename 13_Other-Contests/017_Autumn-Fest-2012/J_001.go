package main

import "fmt"

func main() {
	const MOD = 1000003

	var H, T int
	fmt.Scan(&H, &T)

	var inv [100]int
	var finv [100]int
	var dp [100][100][100]int
	inv[1] = 1
	for i := 2; i < 100; i++ {
		inv[i] = inv[MOD%i] * (MOD - MOD/i) % MOD
	}
	finv[0] = 1
	for i := 1; i < 100; i++ {
		finv[i] = finv[i-1] * inv[i] % MOD
	}

	dp[0][0][0] = 1
	for i := 0; i < 80; i++ {
		for j := 0; j < 80; j++ {
			for k := 0; k < 80; k++ {
				if dp[i][j][k] > 0 {
					for k2 := 1; k2 <= 9; k2++ {
						low := max(k2*(k2-1), 1)
						high := k2 * (k2 + 1)
						for j2 := low; j2 <= high; j2++ {
							if j+j2 < 80 {
								dp[i+1][j+j2][k+k2] += dp[i][j][k]
								if dp[i+1][j+j2][k+k2] >= MOD {
									dp[i+1][j+j2][k+k2] -= MOD
								}
							}
						}
					}
				}
			}
		}
	}

	ans := 0
	for i := 0; i < 80; i++ {
		for k := 0; k < 80; k++ {
			sum := 0
			for j := 0; j < H; j++ {
				sum = (sum + dp[i][j][k]) % MOD
			}
			if sum > 0 {
				coef := finv[i+1]
				for j := T - k + i + 1; j > T-k; j-- {
					coef = coef * j % MOD
				}
				ans = (ans + coef*sum) % MOD
			}
		}
	}

	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
