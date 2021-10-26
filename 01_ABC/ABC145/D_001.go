package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	res := 0
	mod := int(1e9 + 7)
	if (x+y)%3 == 0 {
		n := (x + y) / 3
		k := min(x, y) - n
		initMod(mod)
		res = nCrMod(n, k, mod)
	}

	fmt.Println(res)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

const size = 1000000

var fact, invf [size]int

func initMod(mod int) {
	fact[0] = 1
	invf[0] = 1
	for i := int(1); i < size; i++ {
		fact[i] = (fact[i-1] * i) % mod
		invf[i] = invMod(fact[i], mod)
	}
}

func powMod(a, n, mod int) int {
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

func invMod(a, mod int) int {
	return powMod(a, mod-2, mod)
}

func nCrMod(n, r, mod int) int {
	if n < r || n < 0 || r < 0 {
		return 0
	}
	return fact[n] * invf[r] % mod * invf[n-r] % mod
}
