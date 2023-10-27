package main

import "fmt"

func main() {
	var N, K int
	fmt.Scan(&N, &K)
	p := make([]int, N+1)
	for i := range p {
		p[i] = 1
	}
	for i := K + 1; i < N+1; i++ {
		p[i] = powMod(2, (i-K)*(i-K+1)/2)
	}
	dp := make([]int, N+2)
	dp[0] = 1
	for i := 0; i < N+1; i++ {
		for j := i + 1; j < N+2; j++ {
			dp[j] = (dp[j] - dp[i]*p[j-i-1]) % MOD
		}
	}
	fmt.Println((MOD - dp[N+1]) % MOD)
}

const MOD = 1000000007

func powMod(a, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = res * a % MOD
		}
		a = a * a % MOD
		n /= 2
	}
	return res
}
