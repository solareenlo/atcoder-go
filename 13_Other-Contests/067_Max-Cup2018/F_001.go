package main

import "fmt"

func main() {
	const MOD = 1000000007

	var n, K, L int
	fmt.Scan(&n, &K, &L)
	var dp [400][8][8][8][8]int
	dp[0][0][0][0][0] = 1
	for i := 0; i < n; i++ {
		for j := 0; j < 8; j++ {
			for k := 0; k < 8; k++ {
				for t := 0; t < 8; t++ {
					for l := 0; l < 8; l++ {
						if dp[i][j][k][t][l] == 0 {
							continue
						}
						var cnt [3]int
						a := []int{j, k, t, l}
						for p := 0; p < 4; p++ {
							if 4-p > K-1 {
								continue
							}
							for q := 0; q < 3; q++ {
								if ((a[p] >> q) & 1) != 0 {
									cnt[q]++
								}
							}
						}
						for p := 1; p < 8; p++ {
							flag := true
							for q := 0; q < 3; q++ {
								if (p>>q&1) != 0 && cnt[q] == L {
									flag = false
								}
							}
							if !flag {
								continue
							}
							dp[i+1][k][t][l][p] = (dp[i+1][k][t][l][p] + dp[i][j][k][t][l]) % MOD
						}
					}
				}
			}
		}
	}

	ans := 0
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			for k := 0; k < 8; k++ {
				for l := 0; l < 8; l++ {
					ans = (ans + dp[n][i][j][k][l]) % MOD
				}
			}
		}
	}
	fmt.Println(ans)
}
