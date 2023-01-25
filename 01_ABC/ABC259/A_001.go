package main

import "fmt"

func main() {
	var n, m, x, t, d int
	fmt.Scan(&n, &m, &x, &t, &d)
	x = max(m, x)

	fmt.Println(t - (x-m)*d)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
