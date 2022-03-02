package main

import (
	"bufio"
	"fmt"
	"os"
)

const mod = 998244353

var (
	n   int
	dp      = [100005][2]int{}
	a       = [100005]int{}
	iv3 int = (mod + 1) / 3
)

func calc(x, y int) int {
	for i := range dp {
		for j := range dp[i] {
			dp[i][j] = 0
		}
	}
	dp[0][x] = 1
	for i := 0; i < n; i++ {
		g := a[i] - y
		c := (g + 1) * g / 2 % mod
		dp[i+1][0] = (c*dp[i][0] + (g+1)*dp[i][1]) % mod
		g += y
		c = (g + 1) * g / 2 % mod
		ss := c * (g - 1) % mod * iv3 % mod
		dp[i+1][1] = (c*dp[i][1] + ss*dp[i][0]) % mod
	}
	return dp[n][x]
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n)

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	fmt.Println((calc(0, 0) + calc(1, 0) - calc(0, 1) - calc(1, 1) + 2*mod) % mod)
}
