package main

import "fmt"

func main() {
	const N = 1005
	const P = 998244353

	var n, m, k int
	fmt.Scan(&n, &m, &k)

	var inv [N]int
	inv[1] = 1
	for i := 2; i <= m; i++ {
		inv[i] = ((P - P/i) * inv[P%i]) % P
	}

	var f [N][N]int
	f[0][0] = 1
	ans := 0
	for i := 1; i <= k; i++ {
		for j := 0; j <= n; j++ {
			for l := 1; l <= m; l++ {
				pos := j + l
				if pos > n {
					pos = n - (pos - n)
				}
				f[i][pos] = (f[i][pos] + f[i-1][j]*inv[m]) % P
			}
		}
		ans = (ans + f[i][n]) % P
		f[i][n] = 0
	}
	fmt.Println(ans)
}
