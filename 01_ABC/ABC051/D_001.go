package main

import "fmt"

type tuple struct {
	a, b, c int
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	v := make([][]int, n)
	for i := range v {
		v[i] = make([]int, n)
		for j := range v[i] {
			if i != j {
				v[i][j] = int(1e9 + 7)
			}
		}
	}

	e := make([]tuple, m)
	var a, b, c int
	for i := 0; i < m; i++ {
		fmt.Scan(&a, &b, &c)
		a--
		b--
		e[i] = tuple{a: a, b: b, c: c}
		v[a][b] = c
		v[b][a] = c
	}

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				v[i][j] = min(v[i][j], v[i][k]+v[k][j])
			}
		}
	}

	res := 0
	for i := 0; i < m; i++ {
		if v[e[i].a][e[i].b] != e[i].c {
			res++
		}
	}
	fmt.Println(res)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
