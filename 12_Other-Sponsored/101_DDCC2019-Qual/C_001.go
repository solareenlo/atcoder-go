package main

import "fmt"

const MOD = 1000000007

func f(a int) int {
	b := a * a % MOD
	a = b * b % MOD
	a = a * a % MOD
	return a * b % MOD
}

func main() {
	var n int
	fmt.Scan(&n)
	r := 0
	for k := 1; k <= n; k++ {
		r += f(n/k) * (f(k) - f(k-1) + MOD) % MOD
		r %= MOD
	}
	fmt.Println(r)
}
