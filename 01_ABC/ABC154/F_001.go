package main

import "fmt"

func main() {
	var r1, c1, r2, c2 int
	fmt.Scan(&r1, &c1, &r2, &c2)
	r2++
	c2++

	initMod()
	res := nCrMod(r2+c2, r2) - nCrMod(r1+c2, r1) - nCrMod(r2+c1, c1) + nCrMod(r1+c1, r1)
	fmt.Println((res%mod + mod) % mod)
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
