package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	d := make([][]int, n)
	for i := range d {
		d[i] = make([]int, n)
		for j := range d[i] {
			d[i][j] = int(1e9)
		}
		d[i][i] = 0
	}

	var a, b int
	for i := 0; i < m; i++ {
		fmt.Scan(&a, &b)
		a--
		b--
		d[a][b] = 1
		d[b][a] = 1
	}

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				d[i][j] = min(d[i][j], d[i][k]+d[k][j])
			}
		}
	}

	for i := 0; i < n; i++ {
		cnt := 0
		for j := 0; j < n; j++ {
			if d[i][j] == 2 {
				cnt++
			}
		}
		fmt.Println(cnt)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
