package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	dp := make([]int, a)
	dp[0] = 1
	for i := 1; i < a; i++ {
		dp[i] = divMod((dp[i-1] * (c + i - 1) % mod), i)
	}
	for i := 1; i < a; i++ {
		dp[i] = (dp[i] + dp[i-1]*b%mod) % mod
	}
	fmt.Println(dp[a-1])
}

const mod = 1000000007

func divMod(a, b int) int {
	ret := a * modInv(b)
	ret %= mod
	return ret
}

func modInv(a int) int {
	b, u, v := mod, 1, 0
	for b != 0 {
		t := a / b
		a -= t * b
		a, b = b, a
		u -= t * v
		u, v = v, u
	}
	u %= mod
	if u < 0 {
		u += mod
	}
	return u
}
