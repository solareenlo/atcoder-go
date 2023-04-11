package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const mod = 1000000007

	var n, m int
	fmt.Fscan(in, &n, &m)

	var dp [1005][1005]int
	dp[1][1] = 1
	for i := 1; i <= n; i++ {
		var s string
		fmt.Fscan(in, &s)
		s = " " + s
		for j := 1; j <= m; j++ {
			if (i != 1 || j != 1) && (s[j] != '#') {
				dp[i][j] = (dp[i-1][j] + dp[i][j-1]) % mod
			}
		}
	}
	fmt.Println(dp[n][m])
}
