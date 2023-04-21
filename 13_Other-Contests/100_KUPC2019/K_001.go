package main

import "fmt"

func main() {
	initMod()

	var n, m, p, q, r int
	fmt.Scan(&n, &m, &p, &q, &r)
	p = (p + n) % m
	q = (q + n) % m
	r = (r + n) % m
	var coef [size]int
	for i := 0; i < n+1; i++ {
		coef[i%m] = (coef[i%m] + nCrMod(n, i)) % mod
	}
	ret := 0
	for t := 0; t < 2; t++ {
		if (p+q+r)%2 == 0 {
			a := ((p+q+r)/2 - p + m) % m
			b := ((p+q+r)/2 - q + m) % m
			c := ((p+q+r)/2 - r + m) % m
			ret = (ret + coef[a]*coef[b]%mod*coef[c]%mod) % mod
		}
		p += m
		q += m
		r += m
	}
	fmt.Println(ret)
}

const mod = 998244353
const size = 1000005

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
