package main

import "fmt"

const mod = 1000000007

var N, M, L int

func f(L int) int {
	dp := [303][303]int{}
	dp[0][0] = 1

	for i := 0; i < N; i++ {
		for j := 0; j <= i; j++ {
			T := dp[i][j]
			for k := 1; i+k <= N && j+k-1 <= M && k <= L; k++ {
				if k > 1 {
					dp[i+k][j+k] += T
					dp[i+k][j+k] %= mod
				}
				if k == 2 {
					T *= 500000004
					T %= mod
				}
				dp[i+k][j+k-1] += T * k % mod
				dp[i+k][j+k-1] %= mod
				T *= N - i - k
				T %= mod
			}
		}
	}

	return dp[N][M]
}

func main() {
	fmt.Scan(&N, &M, &L)
	fmt.Println((f(L) - f(L-1) + mod) % mod)
}
