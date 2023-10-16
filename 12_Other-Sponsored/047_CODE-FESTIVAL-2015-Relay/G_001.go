package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m, L int
	fmt.Fscan(in, &n, &m, &L)

	var c1, w1, c2, w2, dp [10007]int
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &c1[i], &w1[i])
	}
	for i := 1; i <= m; i++ {
		fmt.Fscan(in, &c2[i], &w2[i])
	}
	for i := 1; i <= m; i++ {
		for j := L; j >= c2[i]; j-- {
			dp[j] = max(dp[j], dp[j-c2[i]]+w2[i])
		}
	}
	ans := -1
	for i := 1; i <= n; i++ {
		if c1[i] <= L {
			ans = max(ans, dp[L-c1[i]]+w1[i])
		}
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
