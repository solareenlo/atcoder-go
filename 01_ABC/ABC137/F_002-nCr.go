package main

import "fmt"

func main() {
	var p int
	fmt.Scan(&p)

	a := make([]int, p)
	for i := 0; i < p; i++ {
		fmt.Scan(&a[i])
	}

	initMod(p)
	b := make([]int, p)
	for j := 0; j < p; j++ {
		if a[j] != 0 {
			b[0] = (b[0] + 1) % p
			tmp := 1
			for i := p - 1; i >= 0; i-- {
				b[i] = (b[i] - nCrMod(p-1, i, p)*tmp) % p
				if b[i] < 0 {
					b[i] += p
				}
				tmp = (tmp * (p - j)) % p
			}
		}
	}

	for i := 0; i < p; i++ {
		fmt.Print(b[i])
		if i != p-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}

const size = 202020

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
