package main

import "fmt"

func main() {

	initMod()

	var n, K int
	var s string
	fmt.Scan(&n, &K, &s)

	c1, c2 := 0, 0
	for i := 0; i < n/2; i++ {
		if s[i] != s[len(s)-i-1] {
			K--
			c1++
		} else {
			c2++
		}
	}
	if K < 0 {
		fmt.Println(0)
		return
	}
	var sum [300000]int
	for i := 0; i <= c1; i++ {
		if i != 0 {
			sum[i] = (sum[i] + sum[i-1]) % mod
		}
		sum[i] = (sum[i] + (nCrMod(c1, i)*powMod(24, i)%mod)*powMod(2, c1-i)%mod) % mod
	}
	ans := 0
	for i := 0; i <= min(c2, K/2); i++ {
		D := nCrMod(c2, i) * powMod(25, i) % mod
		if n%2 == 1 && 2*i+1 <= K {
			ans = (ans + (D*25%mod)*sum[min(c1, K-(2*i+1))]%mod) % mod
		}
		ans = (ans + D*sum[min(c1, K-2*i)]%mod) % mod
	}
	if c1 == 0 && K == 1 {
		ans = (ans + mod - 1) % mod
	}
	fmt.Println(ans)
}

const mod = 1000000007
const size = 200005

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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
