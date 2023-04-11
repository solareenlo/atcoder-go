package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	const MAXV = 100000
	var dp [MAXV + 5]int
	for i := 1; i < MAXV+1; i++ {
		dp[i] = 1e9 + 1
	}

	var n, m int
	fmt.Fscan(in, &n, &m)

	for i := 1; i <= n; i++ {
		var w, v int
		fmt.Fscan(in, &w, &v)
		for j := MAXV; j >= v; j-- {
			dp[j] = min(dp[j], dp[j-v]+w)
		}
	}
	for i := MAXV; i > 0; i-- {
		if dp[i] <= m {
			fmt.Fprintln(out, i)
			break
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
