package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	tot := powMod(2, n)
	ans := 1
	prd := tot
	tmp := 1
	for i := 0; i <= n; i++ {
		ans = (ans + prd) % mod
		prd = (prd * divMod((tot-tmp+mod)%mod, i+2)) % mod
		tmp = (tmp * 2) % mod
	}
	fmt.Println(ans)
}

const mod = 998244353

func powMod(a, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = res * a % mod
		}
		a = a * a % mod
		n /= 2
	}
	return res
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
