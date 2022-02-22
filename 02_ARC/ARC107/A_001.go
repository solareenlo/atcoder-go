package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	const mod = 998244353
	sumB := divMod((b * (1 + b) % mod), 2)
	sumA := divMod((a * (1 + a) % mod), 2)
	sumC := divMod((c * (1 + c) % mod), 2)
	fmt.Println(((sumA * sumB % mod) * sumC) % mod)
}

const mod = 998244353

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
