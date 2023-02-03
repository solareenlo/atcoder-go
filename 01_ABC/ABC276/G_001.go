package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	initMod()
	ans := 0
	for c1 := 0; c1 <= n+1; c1++ {
		for c2 := max(0, n-c1-1); c2 <= n-c1+1; c2++ {
			c3 := (m - c1 - c2*2) / 3
			if c3 < 0 || (m-c1-c2*2)%3 != 0 {
				continue
			}
			tmp := 0
			if c1+c2 == n {
				tmp++
			}
			ans += (1 + tmp) * nCrMod(c3+n, n) * nCrMod(c1+c2, c1) % mod
			ans %= mod
		}
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

const mod = 998244353
const size = 10000010

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
