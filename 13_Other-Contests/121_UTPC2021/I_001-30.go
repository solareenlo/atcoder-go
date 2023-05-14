package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const INF = 2147483647

	const N = 22
	var cost [N][N]int
	var dp [1 << N][N]int
	var a [N]int

	var n, m int
	fmt.Fscan(in, &n, &m)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Fscan(in, &a[j])
			a[j]--
		}
		for j := 0; j < m; j++ {
			for k := 0; k < j; k++ {
				if a[j] >= 0 && a[k] >= 0 {
					cost[a[j]][a[k]]++
				}
			}
		}
	}

	for i := 0; i < N; i++ {
		for j := 0; j < (1 << N); j++ {
			dp[j][i] = INF
		}
	}
	for i := 0; i < m; i++ {
		dp[1<<i][i] = 0
	}

	var x, y [N]int
	xc, yc := 0, 0
	for b := 1; b < (1 << m); b++ {
		xc = 0
		yc = 0
		for i := 0; i < m; i++ {
			if ((b >> i) & 1) != 0 {
				x[xc] = i
				xc++
			} else {
				y[yc] = i
				yc++
			}
		}
		pc := popcount(uint64(b))
		for ii := 0; ii < xc; ii++ {
			i := x[ii]
			for jj := 0; jj < yc; jj++ {
				j := y[jj]
				dp[b|(1<<j)][j] = min(dp[b|(1<<j)][j], dp[b][i]+cost[i][j]*(m-pc))
			}
		}
	}
	mini := INF
	for i := 0; i < m; i++ {
		mini = min(mini, dp[mask(m)][i])
	}
	fmt.Println(mini)
}

func mask(n int) int {
	return (1 << n) - 1
}

func popcount(n uint64) int {
	return bits.OnesCount64(n)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
