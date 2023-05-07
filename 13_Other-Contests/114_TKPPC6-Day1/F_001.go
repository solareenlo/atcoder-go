package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	initMod()

	dp := make([]mint, n+k+1)
	sumv := mint(1)
	dp[0] = 1
	for i := 1; i < n; i++ {
		dp[i] = sumv
		if i >= k {
			sumv = (sumv - mint(2).pow(k-1)*dp[i-k]%MOD + MOD) % MOD
		}
		sumv = (2*sumv%MOD + dp[i]) % MOD
	}
	fmt.Println(dp[n-1])
}

type mint int

func (m mint) pow(p int) mint {
	return powMod(m, p)
}

func (m mint) inv() mint {
	return invMod(m)
}

func (m mint) div(n mint) mint {
	return divMod(m, n)
}

const MOD = 998244353
const VMAX = 1000005

var fact, invf [VMAX]mint

func initMod() {
	fact[0] = 1
	invf[0] = 1
	for i := mint(1); i < VMAX; i++ {
		fact[i] = (fact[i-1] * i) % MOD
		invf[i] = invMod(fact[i])
	}
}

func powMod(a mint, n int) mint {
	res := mint(1)
	for n > 0 {
		if n%2 == 1 {
			res = res * a % MOD
		}
		a = a * a % MOD
		n /= 2
	}
	return res
}

func invMod(a mint) mint {
	return powMod(a, MOD-2)
}

func divMod(a, b mint) mint {
	ret := a * modInv(b)
	ret %= MOD
	return ret
}

func modInv(a mint) mint {
	b, u, v := mint(MOD), mint(1), mint(0)
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
