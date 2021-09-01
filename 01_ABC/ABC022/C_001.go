package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	d := make([][]int, n+1)
	for i := range d {
		d[i] = make([]int, n+1)
		for j := range d[i] {
			if i != j {
				d[i][j] = int(1e9 + 7)
			}
		}
	}

	var u, v, l int
	for i := 0; i < m; i++ {
		fmt.Scan(&u, &v, &l)
		d[u][v] = l
		d[v][u] = l
	}

	for k := 2; k <= n; k++ {
		for i := 2; i <= n; i++ {
			for j := 2; j <= n; j++ {
				d[i][j] = min(d[i][j], d[i][k]+d[k][j])
			}
		}
	}

	res := int(1e9 + 7)
	for i := 2; i <= n; i++ {
		for j := i + 1; j <= n; j++ {
			res = min(res, d[i][j]+d[1][i]+d[1][j])
		}
	}
	if res != int(1e9+7) {
		fmt.Println(res)
	} else {
		fmt.Println(-1)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
