package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	const mod = 998244353
	d := [1010][12][12][12]int{}
	d[0][m+1][m+1][m+1] = 1
	for i := 1; i <= n; i++ {
		for a := 1; a <= m+1; a++ {
			for b := a; b <= m+1; b++ {
				for c := b; c <= m+1; c++ {
					for j := 1; j <= m; j++ {
						tmp := d[i-1][a][b][c]
						if j <= a {
							d[i][j][b][c] += tmp
							d[i][j][b][c] %= mod
						} else if j <= b {
							d[i][a][j][c] += tmp
							d[i][a][j][c] %= mod
						} else if j <= c {
							d[i][a][b][j] += tmp
							d[i][a][b][j] %= mod
						}
					}
				}
			}
		}
	}

	ans := 0
	for a := 1; a <= m; a++ {
		for b := a + 1; b <= m; b++ {
			for c := b + 1; c <= m; c++ {
				ans += d[n][a][b][c]
				ans %= mod
			}
		}
	}
	fmt.Println(ans)
}
