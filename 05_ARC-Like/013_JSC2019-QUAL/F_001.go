package main

import "fmt"

var (
	n int
	m int
)

func dfs(l, r, s int) int {
	if r*m > s {
		return 0
	}
	s -= r * m
	ans := 0
	l++
	for i := 0; i <= n-m && i*l <= s; i++ {
		o := nCrMod(n-m, i) * nCrMod(s-i*l+n, n) % mod
		if i&1 != 0 {
			ans -= o
		} else {
			ans += o
		}
		ans %= mod
	}
	return ans
}

func calc(s int) int {
	ans := 0
	for x := 1; x <= s; x++ {
		ans = (ans + dfs(x-1, x, s) - dfs(x-1, x+1, s)) % mod
	}
	return ans
}

func main() {
	initMod()

	var ls, rs int
	fmt.Scan(&n, &m, &ls, &rs)
	fmt.Println((nCrMod(rs+n, n) - nCrMod(ls-1+n, n) - nCrMod(n, m)*(calc(rs)-calc(ls-1))%mod + mod*10) % mod)
}

const mod = 1000000007
const size = 2000020

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
