package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const MOD = 1000000007

	var pow [101010]int
	pow[0] = 1

	var D, F, T, N int
	fmt.Fscan(in, &D, &F, &T, &N)
	for i := 1; i <= N; i++ {
		pow[i] = (2 * pow[i-1]) % MOD
	}
	X := make([]int, N+3)
	for i := 1; i <= N; i++ {
		fmt.Fscan(in, &X[i])
	}
	X[N+1] = D
	X[N+2] = int(2e9 + 7)
	dp := make([]int, N+3)
	dp[0] = 1
	for i := 0; i < N+2; i++ {
		if i > 0 {
			dp[i] = (dp[i] + dp[i-1]) % MOD
		}
		l := upperBound(X, X[i]+F-T)
		r := upperBound(X, X[i]+F)
		p := (dp[i] * pow[l-i-1]) % MOD
		dp[l] = (dp[l] + p) % MOD
		dp[r] = (dp[r] - p) % MOD
		if i == 0 {
			dp[i] = 0
		}
	}
	dp[0] = 1
	ans := 0
	for i := 0; i < N+1; i++ {
		if D-X[i] <= F-T {
			ans = (ans + dp[i]*pow[N-i]) % MOD
		}
	}
	ans = (ans + dp[N+1]) % MOD
	fmt.Println((ans + MOD) % MOD)
}

func upperBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] > x
	})
	return idx
}
