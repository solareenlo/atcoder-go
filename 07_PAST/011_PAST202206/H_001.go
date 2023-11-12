package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var dp [303][303]int

	var N, A, B int
	fmt.Fscan(in, &N, &A, &B)
	for N > 0 {
		N--
		var w, v int
		fmt.Fscan(in, &w, &v)
		for a := A; a >= 0; a-- {
			for b := B; b >= 0; b-- {
				if a >= w {
					dp[a][b] = max(dp[a][b], dp[a-w][b]+v)
				}
				if b >= w {
					dp[a][b] = max(dp[a][b], dp[a][b-w]+v)
				}
			}
		}
	}
	fmt.Println(dp[A][B])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
