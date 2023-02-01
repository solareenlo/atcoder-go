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

	var a [105]int
	for i := 1; i <= k; i++ {
		fmt.Fscan(in, &a[i])
	}

	var dp [10005]int
	for i := 1; i <= n; i++ {
		for j := 1; j <= k; j++ {
			if i >= a[j] {
				dp[i] = max(dp[i], i-dp[i-a[j]])
			}
		}
	}

	fmt.Println(dp[n])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
