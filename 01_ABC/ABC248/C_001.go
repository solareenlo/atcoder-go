package main

import "fmt"

func main() {
	var n, m, k int
	fmt.Scan(&n, &m, &k)

	const mod = 998244353
	f := make([]int, 5050)
	f[0] = 1
	for i := 1; i <= n; i++ {
		for j := k; j >= 0; j-- {
			for l := 1; l <= m; l++ {
				f[j+l] += f[j] % mod
			}
			f[j] = 0
		}
	}

	ans := 0
	for i := 1; i <= k; i++ {
		ans += f[i]
		ans %= mod
	}
	fmt.Println(ans)
}
