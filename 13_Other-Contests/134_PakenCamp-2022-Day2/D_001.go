package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const MOD = 998244353

	var N, M int
	fmt.Fscan(in, &N, &M)
	L := make([]int, M)
	R := make([]int, M)
	C := make([]int, M)
	for i := 0; i < M; i++ {
		fmt.Fscan(in, &L[i], &R[i], &C[i])
		L[i]--
		R[i]--
	}
	mx0 := make([]int, N)
	mx1 := make([]int, N)
	for i := 0; i < N; i++ {
		mx0[i] = -1
		mx1[i] = -1
	}
	for i := 0; i < M; i++ {
		if C[i] == 0 {
			mx0[R[i]] = max(mx0[R[i]], L[i])
		}
		if C[i] == 1 {
			mx1[R[i]] = max(mx1[R[i]], L[i])
		}
	}
	dp := make([][]int, 2)
	for i := range dp {
		dp[i] = make([]int, N)
	}
	if mx1[0] == -1 {
		dp[0][0] = 1
	}
	if mx0[0] == -1 {
		dp[1][0] = 1
	}
	sum0 := dp[0][0]
	sum1 := dp[1][0]
	p0 := 0
	p1 := 0
	for i := 1; i < N; i++ {
		if i > mx0[i] {
			dp[1][i] = (dp[1][i] + sum0) % MOD
		}
		if i > mx1[i] {
			dp[0][i] = (dp[0][i] + sum1) % MOD
		}
		sum0 = (sum0 + dp[0][i]) % MOD
		sum1 = (sum1 + dp[1][i]) % MOD
		for p0 <= mx1[i] {
			sum0 = (sum0 + MOD - dp[0][p0]) % MOD
			dp[0][p0] = 0
			p0++
		}
		for p1 <= mx0[i] {
			sum1 = (sum1 + MOD - dp[1][p1]) % MOD
			dp[1][p1] = 0
			p1++
		}
	}
	ans := 0
	for i := 0; i < 2; i++ {
		for j := 0; j < N; j++ {
			ans = (ans + dp[i][j]) % MOD
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
