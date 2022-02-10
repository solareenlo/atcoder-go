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

	p := make([]int, n+1)
	for i := 2; i <= n; i++ {
		fmt.Fscan(in, &p[i])
	}
	x := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &x[i])
	}

	const N = 5010
	dp := [N / 5][N]int{}
	for i := n; i >= 1; i-- {
		for j := x[p[i]]; j >= 0; j-- {
			tmp1 := 1 << 60
			if x[i] <= j {
				tmp1 = dp[p[i]][j-x[i]] + dp[i][x[i]]
			}
			tmp2 := 1 << 60
			if dp[i][x[i]] <= j {
				tmp2 = dp[p[i]][j-dp[i][x[i]]] + x[i]
			}
			dp[p[i]][j] = min(tmp1, tmp2)
		}
	}

	if dp[1][x[1]] < 1<<60 {
		fmt.Println("POSSIBLE")
	} else {
		fmt.Println("IMPOSSIBLE")
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
