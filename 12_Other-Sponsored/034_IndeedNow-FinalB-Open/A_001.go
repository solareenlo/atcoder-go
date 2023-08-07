package main

import "fmt"

func main() {
	const MOD = 1000000007

	var a, b int
	fmt.Scan(&a, &b)
	ans := 0
	for i := a; i <= b; i++ {
		ans += (i * ((i * (i + 1) / 2) % MOD)) % MOD
		ans %= MOD
	}
	fmt.Println(ans)
}
