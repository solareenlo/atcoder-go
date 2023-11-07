package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 300300

	var a, b [N]int
	var dp [N][2]int

	var n int
	fmt.Fscan(in, &n)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i], &b[i])
	}
	if b[1] > 0 {
		dp[1][a[1]] = b[1]
	}
	for i := 2; i <= n; i++ {
		if a[i] == 0 {
			dp[i][1] = dp[i-1][1]
			dp[i][0] = max(dp[i-1][1]+b[i], dp[i-1][0]+max(b[i], 0))
		} else {
			dp[i][0] = dp[i-1][0]
			dp[i][1] = max(dp[i-1][0]+b[i], dp[i-1][1])
		}
	}
	fmt.Println(max(dp[n][0], dp[n][1]))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
