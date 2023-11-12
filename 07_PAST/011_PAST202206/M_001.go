package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	v, l := 1, 1
	for i := 0; i < n; i++ {
		l = l * (i + 1) % MOD
		v = v * (2*n - 1 - i) % MOD
	}
	v = v * ((invMod(l) * n % MOD) * invMod(powMod(2, 2*n-2)) % MOD) % MOD
	v = (v - 1 + MOD) % MOD
	v = v * invMod(2) % MOD
	fmt.Println(v)
}

const MOD = 998244353

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
