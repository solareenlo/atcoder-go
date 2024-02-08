package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, c int
	fmt.Fscan(in, &n, &c)

	var A, B [22]int
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &A[i])
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &B[i])
	}

	dp := make([]int, 1<<n)
	for i := range dp {
		dp[i] = int(1e18)
	}
	dp[0] = -c
	for i := 0; i < 1<<n; i++ {
		for j := 0; j < n; j++ {
			x := popcount(uint32(i))
			p := 0
			t := c
			for k := j; k < n; k++ {
				if (i>>k)&1 != 0 {
					break
				}
				t += abs(A[k] - B[x])
				x++
				p |= 1 << k
				dp[i|p] = min(dp[i|p], dp[i]+t)
			}
		}
	}
	fmt.Println(dp[(1<<n)-1])
}

func popcount(n uint32) int {
	return bits.OnesCount32(n)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
