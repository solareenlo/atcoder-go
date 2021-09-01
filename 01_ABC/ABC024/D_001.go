package main

import "fmt"

func powMod(x, n, mod int64) int64 {
	res := int64(1)
	for n > 0 {
		if n%2 == 1 {
			res = res * x % mod
		}
		x = x * x % mod
		n /= 2
	}
	return res
}

func invMod(x, mod int64) int64 {
	return powMod(x, mod-2, mod)
}

func main() {
	var a, b, C int64
	fmt.Scan(&a, &b, &C)

	const mod int64 = 1e9 + 7
	denom := invMod((a*(b+C)%mod-b*C%mod+mod)%mod, mod)
	r := (b*C%mod - a*C%mod + mod) % mod * denom % mod
	c := (b*C%mod - a*b%mod + mod) % mod * denom % mod
	fmt.Println(r, c)
}
