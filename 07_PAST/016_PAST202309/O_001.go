package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	const MOD = 998244353
	const M = 202020

	var ret, dp [M]int
	var C, add [11][M]int
	var P [101]int
	var B [11]int

	var N, K, Q int
	fmt.Fscan(in, &N, &K, &Q)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &ret[i])
	}
	for i := 0; i < K; i++ {
		fmt.Fscan(in, &P[i])
	}

	for i := 0; i < K; i++ {
		C[i][i] = 1
	}
	for i := K; i <= N; i++ {
		for j := 0; j < K; j++ {
			for x := 0; x < K; x++ {
				C[j][i] = (C[j][i] + C[j][i-(x+1)]*P[x]) % MOD
			}
		}
	}
	for Q > 0 {
		Q--
		var L, R int
		fmt.Fscan(in, &L, &R)
		L--
		for i := 0; i < K; i++ {
			fmt.Fscan(in, &B[i])
		}
		if R-L <= K {
			for i := 0; i < R-L; i++ {
				ret[L+i] = (ret[L+i] + B[i]) % MOD
			}
		} else {
			for i := 0; i < K; i++ {
				ret[L+i] = (ret[L+i] + B[i]) % MOD
				add[i][L+K] = (add[i][L+K] + B[i]) % MOD
				sum := 0
				for j := 0; j < K; j++ {
					sum = (sum + C[j][R-L-K+i]*B[j]) % MOD
				}
				add[i][R] += MOD - sum
			}
		}
	}
	for i := 0; i < N; i++ {
		if i >= K {
			for j := 0; j < K; j++ {
				dp[i-K+j] = (dp[i-K+j] + add[j][i]) % MOD
			}
			for j := 0; j < K; j++ {
				dp[i] = (dp[i] + dp[i-(j+1)]*P[j]) % MOD
			}
		}
		ret[i] = (ret[i] + dp[i]) % MOD
		fmt.Fprintf(out, "%d ", ret[i])
	}
}
