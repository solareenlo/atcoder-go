package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var h, w int
	fmt.Fscan(in, &h, &w)

	dp := [2000][2000]int{}
	dp[0][0] = 1

	x := [2000][2000]int{}
	y := [2000][2000]int{}
	z := [2000][2000]int{}
	mod := int(1e9 + 7)
	for i := 0; i < h; i++ {
		var c string
		fmt.Fscan(in, &c)
		for j := 0; j < w; j++ {
			if (i == 0 && j == 0) || c[j] == '#' {
				continue
			}
			if j > 0 {
				x[i][j] = (x[i][j-1] + dp[i][j-1]) % mod
			}
			if i > 0 {
				y[i][j] = (y[i-1][j] + dp[i-1][j]) % mod
			}
			if i > 0 && j > 0 {
				z[i][j] = (z[i-1][j-1] + dp[i-1][j-1]) % mod
			}
			dp[i][j] = (x[i][j] + y[i][j] + z[i][j]) % mod
		}
	}

	fmt.Println(dp[h-1][w-1])
}
