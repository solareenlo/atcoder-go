package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	initMod()
	d := gcd(n, m)
	ans := 0
	for i := 0; i < d+1; i++ {
		if gcd(i, n) == 1 && gcd(d-i, m) == 1 {
			ans = (ans + nCrMod(d, i)) % mod
		}
	}
	fmt.Println(ans)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

const mod = 1000000007
const size = 1000010

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
