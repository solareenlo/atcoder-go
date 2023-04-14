package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	initMod()
	ans := nCrMod(n+m-2, n-1)
	l := min(n, m)
	for i := (n + 1) / 2; i <= l; i++ {
		ans -= nCrMod(m-1, i-1) * nCrMod(i, n-i) % mod
	}
	ans %= mod
	if ans < 0 {
		ans += mod
	}
	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

const mod = 1000000007
const size = 200005

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
