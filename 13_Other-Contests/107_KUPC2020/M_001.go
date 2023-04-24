package main

import "fmt"

func main() {
	var n, m, k int
	fmt.Scan(&n, &m, &k)

	initMod()

	pw := 1
	ans := 0
	for i := 0; i <= min(n, m/k); i++ {
		var tmp int
		if i&1 != 0 {
			tmp = (nCrMod(n, i) * (mod - 1) % mod) * pw % mod
		} else {
			tmp = nCrMod(n, i) * pw % mod
		}
		pw = pw * Catp(k, 1) % mod
		ans = (ans + tmp*Catp(m-i*k, n-i)) % mod
	}
	fmt.Println(ans)
}

func Catp(n, m int) int {
	if n == 0 && m == 0 {
		return 1
	}
	return nCrMod(2*n+m-1, n) + mod - nCrMod(2*n+m-1, n+m)
}

const mod = 998244353
const size = 3000010

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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
