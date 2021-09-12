package main

import (
	"fmt"
)

func main() {
	var h, w, a, b int64
	fmt.Scan(&h, &w, &a, &b)

	mod := int64(1e9 + 7)
	res := int64(0)
	initMod(mod)
	r1 := h - 1 - a
	for i := b; i < w; i++ {
		x := nCrMod(i+r1, r1, mod)
		r2 := w - 1 - i
		y := nCrMod((a-1)+r2, r2, mod)
		res += x * y % mod
		res %= mod
	}
	fmt.Println(res)
}

var fact, invf [202020]int64

func initMod(mod int64) {
	fact[0] = 1
	invf[0] = 1
	for i := int64(1); i < 202020; i++ {
		fact[i] = (fact[i-1] * i) % mod
		invf[i] = invMod(fact[i], mod)
	}
}

func powMod(a, n, mod int64) int64 {
	res := int64(1)
	for n > 0 {
		if n%2 == 1 {
			res = res * a % mod
		}
		a = a * a % mod
		n /= 2
	}
	return res
}

func invMod(a, mod int64) int64 {
	return powMod(a, mod-2, mod)
}

func nCrMod(n, r, mod int64) int64 {
	if n < r || n < 0 || r < 0 {
		return 0
	}
	return fact[n] * invf[r] % mod * invf[n-r] % mod
}
