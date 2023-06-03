package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(12 * powMod(7, n-1) % MOD)
}

const MOD = 1000000007

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
