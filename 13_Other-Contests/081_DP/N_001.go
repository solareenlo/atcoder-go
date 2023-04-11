package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	var a, sum [444]int
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		sum[i+1] = sum[i] + a[i]
	}
	var dp [444][444]int
	for i := n; i >= 0; i-- {
		for j := i + 2; j <= n; j++ {
			dp[i][j] = 9e18
			for k := i; k < j; k++ {
				dp[i][j] = min(dp[i][j], dp[i][k]+dp[k][j]+sum[j]-sum[i])
			}
		}
	}
	fmt.Println(dp[0][n])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
