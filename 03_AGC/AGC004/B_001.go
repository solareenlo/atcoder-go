package main

import "fmt"

func main() {
	var n, x int
	fmt.Scan(&n, &x)

	m := [2000][2000]int{}
	for i := 0; i < n; i++ {
		fmt.Scan(&m[0][i])
	}

	a := 1 << 60
	for i := 0; i < n; i++ {
		s := 0
		for j := 0; j < n; j++ {
			if i != 0 {
				m[i][j] = min(m[i-1][j], m[0][(i+j)%n])
			}
			s += m[i][j]
		}
		a = min(a, s+i*x)
	}
	fmt.Println(a)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
