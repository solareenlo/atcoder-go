package main

import "fmt"

func main() {
	const mod = 998244353

	var n, k, c int
	fmt.Scan(&n, &k, &c)

	var f [1000001][2]int
	f[1][0] = c
	for i := 2; i <= n; i++ {
		f[i][0] = f[i-1][0]
		f[i][1] = (f[i][1] + f[i-1][0]*(c-1)) % mod
		f[i][1] = (f[i][1] + f[i-1][1]*2) % mod
		if i > k-1 {
			f[i][1] = (f[i][1] + (f[i-(k-1)][1] * (c - 2) % mod)) % mod
		}
	}
	fmt.Println((f[n][1] + f[n][0]) % mod)
}
