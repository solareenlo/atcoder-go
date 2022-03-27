package main

import "fmt"

func main() {
	var n, a, b, k int
	fmt.Scan(&n, &a, &b, &k)

	initMod()

	ans := 0
	for i := 0; i <= n && a*i <= k; i++ {
		if (k-a*i)%b == 0 {
			ans += nCrMod(n, i) * nCrMod(n, (k-a*i)/b) % mod
			ans %= mod
		}
	}

	fmt.Println(ans)
}

const mod = 998244353
const size = 300005

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
