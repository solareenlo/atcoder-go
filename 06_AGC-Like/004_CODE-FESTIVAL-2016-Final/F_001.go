package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	xo := make([][]int, n+2)
	for i := range xo {
		xo[i] = make([]int, n+2)
	}
	xo[1][0] = 1
	const mod = 1_000_000_007
	for k := 0; k < m; k++ {
		x := make([][]int, n+2)
		for i := range x {
			x[i] = make([]int, n+2)
		}
		for i := 0; i <= n; i++ {
			for j := 0; j <= n; j++ {
				if xo[i][j] != 0 {
					x[i+j][0] = (x[i+j][0] + i*xo[i][j]) % mod
					x[i][j] = (x[i][j] + j*xo[i][j]) % mod
					x[i][j+1] = (x[i][j+1] + (n-i-j)*xo[i][j]) % mod
				}
			}
		}
		x, xo = xo, x
	}
	fmt.Println(xo[n][0])
}
