package main

import "fmt"

func main() {
	var n, a, b, c, d int
	fmt.Scan(&n, &a, &b)

	for i := 0; i < n-1; i++ {
		fmt.Scan(&c, &d)
		x := max((a-1)/c+1, (b-1)/d+1)
		a, b = c*x, d*x
	}
	fmt.Println(a + b)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
