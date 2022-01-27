package main

import "fmt"

func main() {
	var a, b, x, y int
	fmt.Scan(&a, &b, &x, &y)

	l, r := 0, min(a, b)
	for l < r {
		m := (l + r + 1) / 2
		if (a-m)/(x-1)+(b-m)/(y-1) >= m {
			l = m
		} else {
			r = m - 1
		}
	}
	fmt.Println(l)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
