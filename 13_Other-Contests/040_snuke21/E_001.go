package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	initMod()
	res := 0
	for i := 1; i <= n; i++ {
		tmp1 := n * (n - 1) / 2
		tmp2 := i * (n - i)
		tmp3 := nCrMod(n, i)
		tmp4 := powMod(2, tmp1-tmp2)
		res = (res + tmp3*tmp4%mod) % mod
	}
	fmt.Println(res)
}

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

const mod = 1000000007

const size = 101010

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
