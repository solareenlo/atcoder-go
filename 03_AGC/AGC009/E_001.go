package main

import "fmt"

func main() {
	var n, m, k int
	fmt.Scan(&n, &m, &k)

	const mod = 1_000_000_007
	f := [4001][4001]int{}
	f[0][1] = 1
	ans := 0
	for i := 1; i < (n+m-1)/(k-1)+1; i++ {
		for j := 0; j < n+1; j++ {
			for I := 0; I < k; I++ {
				f[i][j+I] += f[i-1][j]
				f[i][j+I] %= mod
			}
		}
		ans += f[i][n]
		ans %= mod
	}

	if n > 1 && (n-1)%(k-1) == 0 {
		ans = (ans + mod - 1) % mod
	}
	fmt.Println(ans)
}
