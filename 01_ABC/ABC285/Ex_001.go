package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const MAXN = 10005

	var E, C, dp [MAXN]int

	var n, k int
	fmt.Fscan(in, &n, &k)
	for i := 1; i <= k; i++ {
		fmt.Fscan(in, &E[i])
	}
	C[0] = 1
	for i := 1; i <= n; i++ {
		C[i] = divMod(C[i-1]*(n-i+1)%MOD, i)
	}
	dp[0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j < MAXN; j++ {
			dp[j] = (dp[j] + dp[j-1]) % MOD
		}
	}
	res := 0
	for i := 0; i < n+1; i++ {
		cur := 1
		for j := 1; j <= k; j++ {
			cur = cur * dp[E[j]] % MOD
		}
		if (i & 1) != 0 {
			res = (res - C[i]*cur%MOD + MOD) % MOD
		} else {
			res = (res + C[i]*cur%MOD) % MOD
		}
		for j := MAXN - 1; j >= 1; j-- {
			dp[j] = (dp[j] - dp[j-1] + MOD) % MOD
		}
		for j := 2; j < MAXN; j++ {
			dp[j] = (dp[j] + dp[j-2]) % MOD
		}
	}
	fmt.Println(res)
}

const MOD = 1000000007

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
