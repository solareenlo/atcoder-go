package main

import "fmt"

func main() {
	var N, M int
	fmt.Scan(&N, &M)

	mod = M

	comb := make([][]int, N+1)
	for i := range comb {
		comb[i] = make([]int, N+1)
	}
	comb[0][0] = 1

	for i := 0; i < N; i++ {
		for j := 0; j < i+2; j++ {
			if j == 0 {
				comb[i+1][j] = 1
			} else {
				comb[i+1][j] = (comb[i][j-1] + comb[i][j]) % mod
			}
		}
	}

	pw := make([]int, N*N+1)
	for i := range pw {
		pw[i] = 1
	}
	for i := 0; i < N*N; i++ {
		pw[i+1] = (pw[i] * 2) % mod
	}

	cnt := make([][]int, N)
	for i := range cnt {
		cnt[i] = make([]int, N)
	}
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			cnt[i][j] = powMod(pw[i]-1, j)
		}
	}

	dp := make([][]int, N)
	for i := range dp {
		dp[i] = make([]int, N-1)
	}
	dp[0][0] = 1
	for i := 0; i < N-2; i++ {
		for j := 0; j < i+1; j++ {
			for k := 0; k < N-2-i; k++ {
				dp[i+k+1][k] = (dp[i+k+1][k] + ((((dp[i][j]*comb[N-2-i][k+1])%mod)*cnt[j+1][k+1])%mod)*pw[k*(k+1)/2]%mod) % mod
			}
		}
	}
	for j := 0; j < N-2; j++ {
		dp[N-1][0] = (dp[N-1][0] + (dp[N-2][j]*(pw[j+1]-1))%mod) % mod
	}
	fmt.Println(dp[N-1][0])
}

var mod int

func powMod(a, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = res * a % mod
		}
		a = a * a % mod
		n /= 2
	}
	return res
}
