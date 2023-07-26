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

	var r [3001]int
	var dp [3000][2]int
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &r[i])
	}
	for i := 0; i < n; i++ {
		dp[i][0] = 1
		dp[i][1] = 1
		for j := 0; j < i; j++ {
			if r[j] < r[i] {
				dp[i][0] = max(dp[i][0], dp[j][1]+1)
			}
			if r[j] > r[i] {
				dp[i][1] = max(dp[i][1], dp[j][0]+1)
			}
		}
	}
	ans := max(dp[n-1][0], dp[n-1][1])
	if ans < 3 {
		ans = 0
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
