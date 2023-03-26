package main

import (
	"fmt"
)

const MOD = 1e9 + 7

func main() {
	var s string
	var d int
	fmt.Scan(&s, &d)

	dp := make([][]int64, len(s)+1)
	for i := 0; i <= len(s); i++ {
		dp[i] = make([]int64, d+1)
	}
	dp[0][0] = 1

	for i := 0; i < len(s); i++ {
		for j := 0; j <= d; j++ {
			if dp[i][j] == 0 {
				continue
			}

			sum := int64(s[i] - '0')
			for k := i + 1; k <= len(s) && j+int(sum) <= d; k++ {
				dp[k][j+int(sum)] += dp[i][j]
				dp[k][j+int(sum)] %= MOD

				if k < len(s) {
					sum = sum*10 + int64(s[k]-'0')
				}
			}
		}
	}

	var res int64
	for i := 0; i <= d; i++ {
		res += dp[len(s)][i]
		res %= MOD
	}

	fmt.Println(res)
}
