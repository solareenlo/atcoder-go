package main

import "fmt"

func main() {
	const MOD = 998244353
	const N = 2000200

	var f [N]int

	var n, m int
	fmt.Scan(&n, &m)
	f[1] = m
	f[2] = m * (m - 1) % MOD
	f[3] = ((m * (m - 1) % MOD) * (m - 2) % MOD)
	for i := 4; i <= n; i++ {
		f[i] = ((m-2)*f[i-1]%MOD + (m-1)*f[i-2]%MOD) % MOD
	}
	fmt.Println(f[n])
}
