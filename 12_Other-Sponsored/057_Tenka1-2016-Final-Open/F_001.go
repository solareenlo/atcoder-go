package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const MX = 45
	const maxv = 45

	var n, G int
	fmt.Fscan(in, &n, &G)
	ans := 0.0
	var a [MX]int
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		ans += float64(a[i])
	}
	var dp [MX][MX * MX]float64
	dp[0][0] = 1
	for i := 1; i <= n; i++ {
		for x := i - 1; x >= 0; x-- {
			for y := 0; y <= x*maxv; y++ {
				if dp[x][y] == 0.0 {
					continue
				}
				dp[x+1][y+a[i]] += dp[x][y] / float64(n-x) * float64(x+1)
			}
		}
	}
	for i := 1; i <= n; i++ {
		for j := G * n; j < n*maxv; j++ {
			if dp[i][j] == 0.0 {
				continue
			}
			ans -= dp[i][j] * float64(j-G*n) / float64(i)
		}
	}
	fmt.Printf("%.10f\n", ans)
}
