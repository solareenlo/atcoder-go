package main

import "fmt"

func main() {
	var n, m, x int
	fmt.Scan(&n, &m, &x)

	c := make([]int, n)
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&c[i])
		a[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&a[i][j])
		}
	}

	res := 1 << 60
	for bit := 0; bit < 1<<n; bit++ {
		level := make([]int, m)
		cost := 0
		for i := 0; i < n; i++ {
			if (bit>>i)&1 != 0 {
				cost += c[i]
				for j := 0; j < m; j++ {
					level[j] += a[i][j]
				}
			}
		}
		ok := true
		for j := 0; j < m; j++ {
			if level[j] < x {
				ok = false
			}
		}
		if ok {
			res = min(res, cost)
		}
	}

	if res == 1<<60 {
		fmt.Println(-1)
	} else {
		fmt.Println(res)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
