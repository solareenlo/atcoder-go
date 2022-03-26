package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	n := len(s)
	s = " " + s

	const mod = 1_000_000_007
	g := [2][7]int{{1, 3, 4, 1, 6, 5, 4}, {2, 0, 5, 1, 2, 5, 4}}

	f := [300005][7]int{}
	f[0][0] = 1
	for i := 0; i < n; i++ {
		for j := 0; j < 7; j++ {
			if s[i+1] != '1' {
				f[i+1][g[0][j]] += f[i][j]
				f[i+1][g[0][j]] %= mod
			}
			if s[i+1] != '0' {
				f[i+1][g[1][j]] += f[i][j]
				f[i+1][g[1][j]] %= mod
			}
		}
	}
	fmt.Println((f[n][2] + f[n][5]) % mod)
}
