package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	a := make([]int, m)
	b := make([]int, m)
	t := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&a[i], &b[i], &t[i])
		a[i]--
		b[i]--
	}

	d := make([][]int, n)
	for i := range d {
		d[i] = make([]int, n)
		for j := range d[i] {
			d[i][j] = int(1e9)
		}
		d[i][i] = 0
	}

	for i := 0; i < m; i++ {
		d[a[i]][b[i]] = t[i]
		d[b[i]][a[i]] = t[i]
	}

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				d[i][j] = min(d[i][j], d[i][k]+d[k][j])
			}
		}
	}

	res := int(1e9)
	for i := 0; i < n; i++ {
		maxi := 0
		for j := 0; j < n; j++ {
			maxi = max(maxi, d[i][j])
		}
		res = min(res, maxi)
	}
	fmt.Println(res)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
