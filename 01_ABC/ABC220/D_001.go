package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscan(in, &n, &k)

	dp := [100001][10]int{}
	dp[1][k] = 1

	mod := 998244353
	for i := 2; i <= n; i++ {
		var a int
		fmt.Fscan(in, &a)
		for j := 0; j <= 9; j++ {
			dp[i][(j+a)%10] = (dp[i][(j+a)%10] + dp[i-1][j]) % mod
			dp[i][(j*a)%10] = (dp[i][(j*a)%10] + dp[i-1][j]) % mod
		}
	}

	for i := 0; i < 10; i++ {
		fmt.Println(dp[n][i])
	}
}
