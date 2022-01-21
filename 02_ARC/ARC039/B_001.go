package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	initMod()
	res := 0
	if n <= k {
		rem := k % n
		res = nCrMod(n, rem)
	} else {
		res = nCrMod(n+k-1, k)
	}

	fmt.Println(res)
}

const mod = 1000000007
const size = 606

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
