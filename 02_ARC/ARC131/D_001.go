package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m, d int
	fmt.Fscan(in, &n, &m, &d)

	r := make([]int, m+1)
	for i := 0; i < m+1; i++ {
		fmt.Fscan(in, &r[i])
	}
	s := make([]int, m+1)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &s[i])
	}

	dp := make([]int, d)
	for i := 0; i < m; i++ {
		dp[0] += (s[i] - s[i+1]) * min(n/2, r[i+1]/d)
		if r[i+1] < n/2*d && r[i+1]%d != 0 {
			dp[d-r[i+1]%d] += s[i] - s[i+1]
		}
		dp[0] += (s[i] - s[i+1]) * min((n+1)/2, r[i+1]/d+1)
		if r[i+1] < (n+1)/2*d && r[i+1]%d != d-1 {
			dp[r[i+1]%d+1] += s[i+1] - s[i]
		}
	}

	max_sum := dp[0]
	sum := dp[0]
	for i := 0; i < d-1; i++ {
		sum += dp[i+1]
		max_sum = max(max_sum, sum)
	}
	fmt.Println(max_sum)
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
