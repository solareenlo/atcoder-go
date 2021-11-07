package main

import "fmt"

func main() {
	var k int
	var s string
	fmt.Scan(&k, &s)

	n := len(s)

	initMod()

	res := 0
	for i := 0; i < k+1; i++ {
		res += nCrMod(n-1+i, n-1) * powMod(25, i) % mod * powMod(26, k-i) % mod
		res %= mod
	}
	fmt.Println(res)
}

const mod = 1000000007
const size = 2000001

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
