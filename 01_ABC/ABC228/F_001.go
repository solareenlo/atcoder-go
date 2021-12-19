package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 3000
	c := [N][N]int{}
	dp := [N][N]int{}
	var n, m, e, f, g, h int
	fmt.Fscan(in, &n, &m, &e, &f, &g, &h)

	if g > e {
		g = e
	}
	if h > f {
		h = f
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			var k int
			fmt.Fscan(in, &k)
			c[i+1][j+1] = k + c[i][j+1] + c[i+1][j] - c[i][j]
		}
	}
	for i := 0; i < n-g+1; i++ {
		for j := 0; j < m-h+1; j++ {
			dp[i][j] = c[i+g][j+h] - c[i+g][j] - c[i][j+h] + c[i][j]
		}
	}
	for i := 0; i < n-g+1; i++ {
		for k := f - h + 1; k > 1; k -= k / 2 {
			for j := 0; j < m-h-k/2+1; j++ {
				dp[i][j] = max(dp[i][j], dp[i][j+k/2])
			}
		}
	}
	for j := 0; j < m-h+1; j++ {
		for k := e - g + 1; k > 1; k -= k / 2 {
			for i := 0; i < n-g-k/2+1; i++ {
				dp[i][j] = max(dp[i][j], dp[i+k/2][j])
			}
		}
	}
	z := 0
	for i := 0; i < n-e+1; i++ {
		for j := 0; j < m-f+1; j++ {
			z = max(z, c[i+e][j+f]-c[i][j+f]-c[i+e][j]+c[i][j]-dp[i][j])
		}
	}
	fmt.Println(z)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
