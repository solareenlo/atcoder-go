package main

import "fmt"

func main() {
	var n, m, a, b int
	fmt.Scan(&n, &m, &a, &b)

	initMod()

	ans := 0
	for i := 1; i <= n-a; i++ {
		ans = (ans + nCrMod(b+i-2, i-1)*nCrMod(n+m-b-i-1, n-i)%mod) % mod
	}

	fmt.Println(ans)
}

const mod = 1000000007
const size = 233333

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
