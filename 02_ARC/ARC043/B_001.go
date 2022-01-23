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

	dp := make([]int, 100005)
	a := make([]int, 100005)
	m := 0
	for i := 0; i < n; i++ {
		var t int
		fmt.Fscan(in, &t)
		dp[t]++
		a[t]++
		m = max(t, m)
	}

	b := make([]int, 100005)
	for i := 0; i < 3; i++ {
		for j := 1; j < m+1; j++ {
			b[j] = b[j-1] + dp[j]
			dp[j] = a[j] * b[j/2]
		}
	}

	sum := 0
	for i := 0; i < m+1; i++ {
		sum = (sum + dp[i]) % 1000000007
	}
	fmt.Println(sum)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
