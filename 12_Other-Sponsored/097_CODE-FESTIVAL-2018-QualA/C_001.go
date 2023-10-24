package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const MOD = 1000000007

	var dp [2][3005]int

	var n, K int
	fmt.Fscan(in, &n, &K)
	K = min(3000, K)
	dp[0][0] = 1
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(in, &a)
		var t int
		for t = 0; a > 0; a, t = a>>1, t+1 {

		}
		var tmp [2][3005]int
		for i := 0; i < 3001; i++ {
			for j := 0; j < t+1; j++ {
				for k := 0; k < 2; k++ {
					tmp1 := 0
					if k == 0 && j == t {
						tmp1 = 1
					}
					tmp[k+tmp1][min(3000, i+j)] += dp[k][i]
				}
			}
		}
		for i := 0; i < 3001; i++ {
			for j := 0; j < 2; j++ {
				dp[j][i] = tmp[j][i] % MOD
			}
		}
	}
	ans := 0
	for i := 0; i < K; i++ {
		ans += dp[1][i]
	}
	for k := 0; k < 2; k++ {
		ans += dp[k][K]
	}
	fmt.Println(ans % MOD)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
