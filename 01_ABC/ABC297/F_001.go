package main

import "fmt"

func main() {
	initMod()
	var n, m, k int
	fmt.Scan(&n, &m, &k)
	ans := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			ans = (ans + ((C(n*m, k)-(((((C((i-1)*m, k)+C((n-i)*m, k))%MOD)+C((j-1)*n, k))%MOD+C((m-j)*n, k))%MOD+MOD)%MOD)%MOD+((((C((i-1)*(j-1), k)+C((n-i)*(j-1), k))%MOD+C((i-1)*(m-j), k))%MOD+C((n-i)*(m-j), k))%MOD)+MOD)%MOD) % MOD
		}
	}
	fmt.Println(ans * powMod(C(n*m, k), MOD-2) % MOD)
}

const MOD = 998244353
const SIZE = 1010101

var fact, invf [SIZE]int

func initMod() {
	fact[0] = 1
	invf[0] = 1
	for i := 1; i < SIZE; i++ {
		fact[i] = (fact[i-1] * i) % MOD
		invf[i] = invMod(fact[i])
	}
}

func powMod(a, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = res * a % MOD
		}
		a = a * a % MOD
		n /= 2
	}
	return res
}

func invMod(a int) int {
	return powMod(a, MOD-2)
}

func C(n, r int) int {
	if n < r {
		return 0
	}
	if n == r {
		return 1
	}
	if n < 0 || r < 0 {
		return 0
	}
	return (fact[n] * invf[r] % MOD) * invf[n-r] % MOD
}
