package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	a := make([]int, n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i], &b[i])
	}
	c := make([]int, m)
	d := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&c[i], &d[i])
	}

	for i := 0; i < n; i++ {
		res := 0
		mini := int(2e9)
		for j := 0; j < m; j++ {
			dist := abs(a[i]-c[j]) + abs(b[i]-d[j])
			if mini > dist {
				res = j + 1
			}
			mini = min(mini, dist)
		}
		fmt.Println(res)
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
