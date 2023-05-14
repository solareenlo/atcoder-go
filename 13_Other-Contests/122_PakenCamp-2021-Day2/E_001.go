package main

import "fmt"

func main() {
	var N, M int
	fmt.Scan(&N, &M)
	fmt.Println((M % MOD) * ((M - 1) % MOD) % MOD * powMod((M-2)%MOD, N-2) % MOD)
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
