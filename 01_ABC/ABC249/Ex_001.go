package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	c := make([]int, 2002)
	for i := 1; i <= n; i++ {
		var a int
		fmt.Fscan(in, &a)
		c[a]++
	}

	initMod()
	kp := make([]int, n+1)
	kp[0] = 1
	for i := 1; i <= n; i++ {
		kp[i] = (kp[i-1] << 1) % mod
	}

	dp := make([]int, n+1)
	for i := 0; i < n; i++ {
		s := (dp[i]*n%mod - 1) * kp[i+1] % mod
		for j := 0; j <= i; j++ {
			s -= ((n-i)*nCrMod(i+1, j) + 4*i*nCrMod(i-1, j-1)) % mod * dp[j] % mod
		}
		dp[i+1] = (s%mod + mod) * inv[n-i] % mod
	}

	ans := -dp[n]
	for i := 1; i <= n; i++ {
		ans += dp[c[i]]
	}
	fmt.Println((ans%mod + mod) % mod)
}

const mod = 998244353
const size = 2002

var fact, invf, inv [size + 1]int

func initMod() {
	fact[0] = 1
	invf[0] = 1
	for i := int(1); i < size; i++ {
		fact[i] = (fact[i-1] * i) % mod
		invf[i] = invMod(fact[i])
	}
	inv[1] = 1
	for i := 2; i <= size; i++ {
		inv[i] = (mod - mod/i) * inv[mod%i] % mod
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
