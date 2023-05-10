package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	fmt.Println(powMod(2, N-1))
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
