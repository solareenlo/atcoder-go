package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	const MOD = 1_000_000_007
	fmt.Println((N * (N + 1) / 2 % MOD) * (N * (N + 1) / 2 % MOD) % MOD)
}
