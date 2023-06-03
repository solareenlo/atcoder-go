package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const INF = 1e18

	var n int
	fmt.Fscan(in, &n)
	x := make([]float64, n)
	y := make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x[i], &y[i])
	}

	var f func(int, int) float64
	f = func(i, j int) float64 {
		return math.Sqrt(math.Pow(x[i]-x[j], 2) + math.Pow(y[i]-y[j], 2))
	}

	dp := make([][]float64, 1<<n)
	for i := range dp {
		dp[i] = make([]float64, n)
		for j := range dp[i] {
			dp[i][j] = INF
		}
	}
	for i := 0; i < n; i++ {
		dp[0][0] = 0
	}
	for bit := 0; bit < 1<<n; bit++ {
		for i := 0; i < n; i++ {
			if dp[bit][i] == INF {
				continue
			}
			for j := 0; j < n; j++ {
				dp[bit|(1<<j)][j] = math.Min(dp[bit|(1<<j)][j], dp[bit][i]+f(i, j))
			}
		}
	}

	fmt.Println(dp[(1<<n)-1][0])
}
