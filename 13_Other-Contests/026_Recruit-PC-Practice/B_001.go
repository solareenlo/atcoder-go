package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const MOD = 1000000009

	var T int
	fmt.Fscan(in, &T)
	for T > 0 {
		T--
		var a, b, c, d, e, f int
		fmt.Fscan(in, &a, &b, &c, &e, &d, &f)
		dp := make([][][]int, a+1)
		for i := range dp {
			dp[i] = make([][]int, b+1)
			for j := range dp[i] {
				dp[i][j] = make([]int, 512)
			}
		}
		dp[0][0][0] = 1
		for i := 0; i <= a; i++ {
			for j := 0; j <= b; j++ {
				for k := 0; k < 512; k++ {
					anum, bnum := 0, 0
					for l := 0; l < min(c-1, i+j); l++ {
						if ((k >> l) & 1) == 0 {
							anum++
						}
					}
					for l := 0; l < min(d-1, i+j); l++ {
						if ((k >> l) & 1) != 0 {
							bnum++
						}
					}
					if i < a && anum+1 < e {
						nk := (k << 1) % 512
						dp[i+1][j][nk] += dp[i][j][k]
						dp[i+1][j][nk] %= MOD
					}
					if j < b && bnum+1 < f {
						nk := ((k << 1) | 1) % 512
						dp[i][j+1][nk] += dp[i][j][k]
						dp[i][j+1][nk] %= MOD
					}
				}
			}
		}
		ans := 0
		for i := 0; i < 512; i++ {
			ans += dp[a][b][i]
			ans %= MOD
		}
		fmt.Println(ans)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
