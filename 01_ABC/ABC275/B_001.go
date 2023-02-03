package main

import "fmt"

func main() {
	const mod = 998244353

	var a, b, c, d, e, f int
	fmt.Scan(&a, &b, &c, &d, &e, &f)
	g1 := (((a % mod) * (b % mod) % mod) * (c % mod) % mod) % mod
	g2 := (((d % mod) * (e % mod) % mod) * (f % mod) % mod) % mod
	g := (g1%mod - g2%mod + mod) % mod
	fmt.Println(g)
}
