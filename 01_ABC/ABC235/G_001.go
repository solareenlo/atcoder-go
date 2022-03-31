package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, 3)
	fmt.Scan(&a[0], &a[1], &a[2])
	dp := make([]int, 3)
	dp[0] = 1
	dp[1] = 1
	dp[2] = 1

	initMod()

	ans := 0
	for i := 0; i <= n; i++ {
		cur := nCrMod(n, n-i)
		for j := 0; j < 3; j++ {
			cur *= dp[j]
			cur %= mod
			dp[j] += (dp[j] - nCrMod(i, a[j]) + mod) % mod
			dp[j] %= mod
		}
		ans = (cur - ans + mod) % mod
	}
	fmt.Println(ans)
}

const mod = 998244353
const size = 5000005

var fact, invf [size]int

func initMod() {
	fact[0] = 1
	invf[0] = 1
	for i := int(1); i < size; i++ {
		fact[i] = (fact[i-1] * i) % mod
		invf[i] = invMod(fact[i])
	}
}

func powMod(a, n int) int {
	res := int(1)
	for n > 0 {
		if n%2 == 1 {
			res = res * a % mod
		}
		a = a * a % mod
		n /= 2
	}
	return res
}

func invMod(a int) int {
	return powMod(a, mod-2)
}

func nCrMod(n, r int) int {
	if n < r || n < 0 || r < 0 {
		return 0
	}
	return fact[n] * invf[r] % mod * invf[n-r] % mod
}
