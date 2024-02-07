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
	var d [2005]int
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &d[i])
	}
	var l1, c1, k1, l2, c2, k2 int
	fmt.Fscan(in, &l1, &c1, &k1, &l2, &c2, &k2)
	var dp [2005][2005]int
	ans := int(9e18)
	for i := 1; i <= n; i++ {
		for j := 0; j <= k1; j++ {
			dp[i][j] = int(9e18)
			for k := 0; k <= j; k++ {
				dp[i][j] = min(dp[i][j], dp[i-1][j-k]+max(d[i]-k*l1+l2-1, 0)/l2)
			}
			if i == n && dp[i][j] <= k2 {
				ans = min(ans, j*c1+dp[i][j]*c2)
			}
		}
	}
	if ans == int(9e18) {
		fmt.Println(-1)
	} else {
		fmt.Println(ans)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
