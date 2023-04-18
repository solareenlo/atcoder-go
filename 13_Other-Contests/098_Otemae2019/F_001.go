package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, d int
	fmt.Fscan(in, &n, &d)

	var m, dp [2005]int
	for i := 0; i < d; i++ {
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &m[j+1])
			m[j+1] += m[j]
		}
		dp[0] += m[n]
		for j := 0; j < n; j++ {
			dp[j+1] = min(dp[j], abs(m[n]-m[j+1]*2)+dp[j+1])
		}
	}
	fmt.Println(dp[n])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
