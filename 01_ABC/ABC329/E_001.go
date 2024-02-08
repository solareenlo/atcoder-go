package main

import "fmt"

func main() {
	const N = 200005
	const M = 15

	var n, m int
	var s, t string
	fmt.Scan(&n, &m, &s, &t)

	var f [N][M]int
	f[0][0] = 1
	for i := 0; i < n; i++ {
		for j := 0; j <= m; j++ {
			for e := 0; e < m; e++ {
				if s[i] == t[e] && (e == j || e == 0 || j == m) {
					f[i+1][e+1] |= f[i][j]
				}
			}
		}
	}
	if f[n][m] == 1 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
