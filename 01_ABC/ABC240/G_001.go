package main

import "fmt"

func main() {
	var n, x, y, z int
	fmt.Scan(&n, &x, &y, &z)

	x = abs(x)
	y = abs(y)
	z = abs(z)
	if x+y+z > n || ((x+y+z^n)&1) != 0 {
		fmt.Println(0)
		return
	}

	initMod()

	t := (x + y + z + n) >> 1
	s := n + x + y - t
	res := 0
	for k := x + y; k <= s && k <= t; k++ {
		res = res + invf[t-k]*invf[s-k]%mod*invf[k-x]%mod*invf[k-y]%mod*nCrMod(2*k-x-y, k-x-y)
		res %= mod
	}

	fmt.Println(res * fact[n] % mod)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

const mod = 998244353
const size = 10000001

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
