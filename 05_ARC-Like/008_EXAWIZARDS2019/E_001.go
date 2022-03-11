package main

import "fmt"

func main() {
	initMod()

	var b, w int
	fmt.Scan(&b, &w)

	t := 1
	p, q := 0, 0
	for i := 0; i < b+w; i++ {
		if i >= b {
			p = (p + nCrMod(i-1, b-1)*t%mod) % mod
		}
		if i >= w {
			q = (q + nCrMod(i-1, w-1)*t%mod) % mod
		}
		fmt.Println((1 - p + q + mod) % mod * invf[2] % mod)
		t = t * invf[2] % mod
	}
}

const mod = 1000000007
const size = 2000200

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
