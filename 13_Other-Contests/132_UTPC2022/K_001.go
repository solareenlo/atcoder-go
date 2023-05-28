package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const mod = 998244353

	type pair struct {
		x, y int
	}

	var N, M, K int
	fmt.Fscan(in, &N, &M, &K)
	tmp := make([]pair, K)
	res := make([]int, N)
	for i := range res {
		res[i] = mod
	}
	res2 := make([]int, N)
	for i := range res2 {
		res2[i] = -1
	}
	for i := 0; i < K; i++ {
		fmt.Fscan(in, &tmp[i].x, &tmp[i].y)
		tmp[i].x--
		tmp[i].y--
		if tmp[i].x%2 == 0 {
			res2[tmp[i].x/2] = max(res2[tmp[i].x/2], tmp[i].y/2)
		} else {
			res[tmp[i].x/2] = min(res[tmp[i].x/2], tmp[i].y/2)
		}
	}
	var dp [3060][3060]int
	dp[0][M] = 1
	for i := 0; i < N; i++ {
		sum := make([]int, M+1)
		for j := M; j >= 0; j-- {
			sum[j] = dp[i][j]
		}
		for j := M; j >= 1; j-- {
			sum[j-1] = (sum[j-1] + sum[j]) % mod
		}
		for j := res2[i] + 1; j <= min(M, res[i]); j++ {
			dp[i+1][j] = (dp[i+1][j] + sum[j]) % mod
		}
	}
	ans := 0
	for i := 0; i < M+1; i++ {
		ans = (ans + dp[N][i]) % mod
	}
	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
