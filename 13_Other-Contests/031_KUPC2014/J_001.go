package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const INF = int(1e18)

	var N, M, P int
	fmt.Fscan(in, &N, &M, &P)

	co := make([]int, M+1)
	for i := 1; i <= M; i++ {
		var x int
		fmt.Fscan(in, &x)
		co[i] = co[i-1] + x
	}

	var dp [105][105]int
	for i := 0; i < 105; i++ {
		for j := 0; j < 105; j++ {
			dp[i][j] = INF
		}
	}
	dp[0][0] = 0

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			for k := 0; k <= M && j+k <= N; k++ {
				ne := dp[i][j] + co[k]
				if ne <= (i+1)*P {
					dp[i+1][j+k] = min(dp[i+1][j+k], ne)
				}
			}
		}
		if dp[i+1][N] != INF {
			fmt.Println(i + 1)
			return
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
