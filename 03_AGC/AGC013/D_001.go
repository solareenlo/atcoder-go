package main

import "fmt"

const mod = 1_000_000_007

var f = [2][3030]int{}

func calc(m, n int) int {
	for i := 1; i <= m; i++ {
		f[0][i] = 1
	}
	f[0][0] = 0
	f[0][m+1] = 0
	f[1][0] = 0
	f[1][m+1] = 0
	c := 0
	s := 0
	for i := 0; i < n; i++ {
		tmp := 0
		if c == 0 {
			tmp = 1
		}
		for j := 1; j <= m; j++ {
			f[tmp][j] = (f[c][j-1] + f[c][j]*2 + f[c][j+1]) % mod
		}
		c = tmp
	}
	for i := 1; i <= m; i++ {
		s += f[c][i]
		s %= mod
	}
	return s
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	s := 4 * (calc(n, m-1) + mod - calc(n-1, m-1)) % mod
	fmt.Println(s)
}
