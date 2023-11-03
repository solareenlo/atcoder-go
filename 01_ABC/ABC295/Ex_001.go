package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, M int
	fmt.Fscan(in, &N, &M)
	const mod = 998244353

	grid := make([]string, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &grid[i])
	}
	dp := make([][]int, N+1)
	for i := range dp {
		dp[i] = make([]int, 1<<M)
	}
	dp[0][(1<<M)-1] = 1

	for i := 0; i < N; i++ {
		zero := 0
		one := 0
		for j := 0; j < M; j++ {
			if grid[i][j] == '1' {
				one |= (1 << j)
			}
			if grid[i][j] == '0' {
				zero |= (1 << j)
			}
		}
		if zero == 0 {
			copy(dp[i+1], dp[i])
		}
		for j := M - 1; j >= 0; j-- {
			b := (1 << j)
			fix_one := b - 1
			for k := 0; k < (1 << M); k++ {
				if (k & b) != 0 {
					continue
				}
				dp[i][k] = (dp[i][k] + dp[i][k|b]) % mod
				state := k | fix_one
				if ((state & one) == one) && ((state & zero) == 0) {
					dp[i+1][k] = (dp[i+1][k] + dp[i][k]) % mod
				}
			}
		}
	}
	ans := 0
	for i := 0; i < (1 << M); i++ {
		ans = (ans + dp[N][i]) % mod
	}
	fmt.Println(ans)
}
