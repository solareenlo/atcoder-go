package main

import "fmt"

func main() {
	const MOD = 1000000007

	var l, r int
	fmt.Scan(&l, &r)

	x := 1
	ans := 0
	for i := 1; i < 20; i++ {
		if r < x {
			break
		}
		ans += (max(l, x) + r) % MOD * ((r - max(l, x) + 1) % MOD) % MOD * 500000004 % MOD
		ans %= MOD
		x *= 10
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
