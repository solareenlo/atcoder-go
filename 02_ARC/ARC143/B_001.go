package main

import "fmt"

func main() {
	const mod = 998244353

	var n int
	fmt.Scan(&n)

	x := 1
	for i := 1; i <= n; i++ {
		x = (x * i) % mod
	}
	y := 1
	for i := 1; i < 2*n; i++ {
		y = (y * i) % mod
	}
	ans := ((y-x*x)%mod + mod) % mod
	for i := 2 * n; i <= n*n; i++ {
		ans = (ans * i) % mod
	}
	fmt.Println(ans)
}
