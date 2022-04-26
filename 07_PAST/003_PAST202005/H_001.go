package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, L int
	fmt.Fscan(in, &N, &L)

	ng := make([]bool, 1<<17)
	for i := 0; i < N; i++ {
		var x int
		fmt.Fscan(in, &x)
		ng[x] = true
	}

	var T1, T2, T3 int
	fmt.Fscan(in, &T1, &T2, &T3)
	dp := make([]int, 1<<17)
	for i := 1; i <= L; i++ {
		dp[i] = 1 << 30
	}
	for i := 0; i < L; i++ {
		if ng[i] {
			dp[i] += T3
		}
		dp[i+1] = min(dp[i+1], dp[i]+T1)
		dp[i+2] = min(dp[i+2], dp[i]+T1+T2)
		dp[i+4] = min(dp[i+4], dp[i]+T1+T2*3)
	}

	ans := dp[L]
	for k := max(L-3, 0); k < L; k++ {
		ans = min(ans, dp[k]+T1/2+T2/2*((L-k)*2-1))
	}
	fmt.Println(ans)
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
