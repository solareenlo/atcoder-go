package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var x, y [10005]float64
	var dp [10005][25]float64

	var n int
	fmt.Fscan(in, &n, &x[1], &y[1])
	for i := 2; i <= n; i++ {
		fmt.Fscan(in, &x[i], &y[i])
		for j := 0; j < 20; j++ {
			dp[i][j] = 1e9
			for k := max(1, i-j-1); k < i; k++ {
				dp[i][j] = math.Min(dp[i][j], dp[k][j-(i-k-1)]+math.Sqrt((x[i]-x[k])*(x[i]-x[k])+(y[i]-y[k])*(y[i]-y[k])))
			}
		}
	}
	ans := 1e9
	for i := 0; i < 20; i++ {
		if i != 0 {
			tmp := 1 << (i - 1)
			ans = math.Min(ans, dp[n][i]+float64(tmp))
		} else {
			ans = math.Min(ans, dp[n][i])
		}
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
