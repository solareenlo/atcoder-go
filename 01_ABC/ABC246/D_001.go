package main

import "fmt"

func f(a, b int) int {
	return (a*a*a + a*a*b + a*b*b + b*b*b)
}

func main() {
	var n int
	fmt.Scan(&n)

	x := 1 << 60
	j := 1000000
	for i := 0; i <= 1000000; i++ {
		for f(i, j) >= n && j >= 0 {
			x = min(x, f(i, j))
			j--
		}
	}
	fmt.Println(x)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
