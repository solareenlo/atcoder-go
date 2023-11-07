package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(in, &N)
	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &A[i])
	}
	B := make([]int, N)
	for i := 0; i < N; i++ {
		B[i] = min(A[i], 10)
	}
	dp := make([][]int, N+1)
	for i := range dp {
		dp[i] = make([]int, 2048)
	}
	dp[0][1] = 1
	for i := 0; i < N; i++ {
		for j := 0; j < 2048; j++ {
			for k := 1; k <= B[i]; k++ {
				dp[i+1][(j|(j<<k))%2048] = (dp[i+1][(j|(j<<k))%2048] + dp[i][j]) % MOD
			}
			dp[i+1][j] = (dp[i+1][j] + dp[i][j]*(A[i]-B[i])%MOD) % MOD
		}
	}
	ans := 0
	for i := 1024; i < 2048; i++ {
		ans = (ans + dp[N][i]) % MOD
	}
	for i := 0; i < N; i++ {
		ans = divMod(ans, A[i])
	}
	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

const MOD = 998244353

func divMod(a, b int) int {
	ret := a * modInv(b)
	ret %= MOD
	return ret
}

func modInv(a int) int {
	b, u, v := MOD, 1, 0
	for b != 0 {
		t := a / b
		a -= t * b
		a, b = b, a
		u -= t * v
		u, v = v, u
	}
	u %= MOD
	if u < 0 {
		u += MOD
	}
	return u
}
