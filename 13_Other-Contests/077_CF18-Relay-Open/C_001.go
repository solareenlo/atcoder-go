package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, H int
	fmt.Fscan(in, &N, &H)

	h := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &h[i])
	}

	dp := make([]int, 1<<N)
	dp[0] = 1

	for i := 0; i < 1<<N; i++ {
		var S int
		for j := 0; j < N; j++ {
			if ((i >> j) & 1) != 0 {
				continue
			}
			S += h[j]
			if S <= H {
				dp[i|(1<<j)] += dp[i]
			}
		}
	}
	fmt.Println(dp[1<<N-1])
}
