package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(in, &N)

	var A [10][10]int
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			fmt.Fscan(in, &A[i][j])
		}
	}

	var M int
	fmt.Fscan(in, &M)

	var ng [10][10]bool
	for i := 0; i < M; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		x--
		y--
		ng[x][y] = true
		ng[y][x] = true
	}

	var dp [1 << 10][10]int
	for i := 0; i < 1<<N; i++ {
		for j := 0; j < N; j++ {
			dp[i][j] = int(1e9)
		}
	}

	for i := 0; i < N; i++ {
		dp[1<<i][i] = A[i][0]
	}
	for i := 1; i < 1<<N; i++ {
		for j := 0; j < N; j++ {
			for k := 0; k < N; k++ {
				if !ng[j][k] && ((i>>k)&1) == 0 {
					dp[i|1<<k][k] = min(dp[i|1<<k][k], dp[i][j]+A[k][popcount(uint32(i))])
				}
			}
		}
	}

	ans := int(1e9)
	for k := 0; k < N; k++ {
		ans = min(ans, dp[(1<<N)-1][k])
	}
	if ans == int(1e9) {
		ans = -1
	}
	fmt.Println(ans)
}

func popcount(n uint32) int {
	return bits.OnesCount32(n)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
