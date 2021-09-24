package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	d := make([][]int, n+1)
	for i := range d {
		d[i] = make([]int, n+1)
		for j := range d[i] {
			d[i][j] = -int(1e18)
			d[i][i] = 0
		}
	}

	var a, b, c int
	for i := 0; i < m; i++ {
		fmt.Scan(&a, &b, &c)
		d[a][b] = c
	}

	for k := 1; k <= n; k++ {
		for i := 1; i <= n; i++ {
			for j := 1; j <= n; j++ {
				if d[i][k]+d[k][j] > d[i][j] {
					d[i][j] = d[i][k] + d[k][j]
				}
			}
		}
	}
	if d[1][1] > 0 {
		fmt.Println("inf")
	} else {
		fmt.Println(d[1][n])
	}
}
