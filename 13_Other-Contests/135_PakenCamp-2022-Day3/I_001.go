package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const INF = 100000000000000000
	const K = 18

	var n int
	fmt.Fscan(in, &n)
	var d [300000]int
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(in, &a)
		d[a]++
	}
	for i := 0; i < K; i++ {
		for s := 0; s < (1 << K); s++ {
			if ((s >> i) & 1) != 0 {
				d[s] += d[s^(1<<i)]
			}
		}
	}
	var dp [300000]int
	dp[0] = 0
	for s := 1; s < (1 << K); s++ {
		dp[s] = INF
		for i := 0; i < K; i++ {
			if ((s >> i) & 1) == 0 {
				continue
			}
			dp[s] = min(dp[s], dp[s^(1<<i)]+(d[s]-d[s^(1<<i)])*s)
		}
	}
	fmt.Println(dp[(1<<K)-1])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
