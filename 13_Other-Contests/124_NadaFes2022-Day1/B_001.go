package main

import "fmt"

func main() {
	var N, K int
	fmt.Scan(&N, &K)

	ans := 1
	r := 1
	for i, j := 0, N/2; i < N-K; i, j = i+1, j-1 {
		ans = ans * j % MOD
	}
	for i := 1; i < N-K+1; i++ {
		r = r * i % MOD
	}
	ans = ans * invMod(r) % MOD
	fmt.Println(ans)
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
