package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscan(in, &n, &k)

	dp := make([]int, 1<<17)
	for j := 1; j < 1<<k; j++ {
		dp[j] = 1 << 60
	}

	for i := 0; i < n; i++ {
		var A int
		fmt.Fscan(in, &A)
		A = 1 << (A - 1)
		for j := 1 << k; j >= 0; j-- {
			dp[j|A] = min(dp[j|A], dp[j]+bits.OnesCount(uint(j&-A)))
			pc := bits.OnesCount(uint(j))
			dp[j] += min(pc, k-pc)
		}
	}
	fmt.Println(dp[(1<<k)-1])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
