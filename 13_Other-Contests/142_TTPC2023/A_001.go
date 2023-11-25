package main

import "fmt"

func main() {
	const MOD = 998244353

	var N int
	fmt.Scan(&N)
	ans := 0
	for i := 0; i < N; i++ {
		ans *= 2
		ans = ans % MOD
		ans += i
		ans = ans % MOD
	}
	fmt.Println(ans % MOD)
}
